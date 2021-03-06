package pipl

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// SourceLevel is used internally to represent the possible values
// for show_sources in queries to be submitted: {"all", "matching"/"true", "false"}
type SourceLevel string

// MatchRequirements specifies the conditions for a successful person match in our search.
// This is useful for saving money with the Pipl API, as you only need to pay for the
// data you wanted back. If your search results didn't satisfy the match requirements, then
// no data is returned and you don't pay.
type MatchRequirements string

// SourceCategoryRequirements specifies the data categories that must be included in
// results for a successful match. If there is no data from the requested categories,
// then the results returned are empty and you're not charged.
type SourceCategoryRequirements string

const (
	// PiplAPIEndpoint is where we POST queries to
	PiplAPIEndpoint string = "https://api.pipl.com/search/"

	// ShowSourcesNone specifies that we don't need source info back with search results
	ShowSourcesNone SourceLevel = "false"

	// ShowSourcesAll specifies that we want all source info back with our search results
	ShowSourcesAll SourceLevel = "all"

	// ShowSourcesMatching specifies that we want source info that corresponds to data that satisfies our match requirements
	ShowSourcesMatching SourceLevel = "true"

	// MatchRequirementsNone specifies that we don't have any match requirements for this search
	MatchRequirementsNone MatchRequirements = ""

	// SourceCategoryRequirementsNone specifies that we don't require any specific sources in our results.
	SourceCategoryRequirementsNone SourceCategoryRequirements = ""
)

// Client holds client configuration settings
type Client struct {
	// HTTPClient carries out the POST operations
	HTTPClient *http.Client
	// Parameters contains the search parameters that are submitted with your query,
	// which may affect the data returned
	SearchParameters *SearchParameters
}

// SearchParameters holds options that can affect data returned by a search.
type SearchParameters struct {
	// APIKey is required
	APIKey string

	// MinimumProbability is the mininum acceptable probability for inferred data
	MinimumProbability float32

	// InferPersons specifies whether or not the Pipl should return results inferred by statistical analysis
	InferPersons bool

	// MinimumMatch specifies the minimum match confidence for a possible person to be returned in search results
	MinimumMatch float32

	// ShowSources specifies the level of sources info to return with search results, one of {ShowSourcesMatching, ShowSourcesAll, ShowSourcesNone}
	ShowSources SourceLevel

	// HideSponsored specifies whether to omit sponsored data from search results
	HideSponsored bool

	// LiveFeeds specifies whether to use live data sources
	LiveFeeds bool

	// MatchRequirements specifies the criteria for a successful Person match.
	// Results that don't fit your match requirements are discarded. If the remaining
	// search results would be empty, you are not charged for the query.
	MatchRequirements MatchRequirements

	// SourceCategoryRequirements specifies the data categories that must be included in
	// results for a successful match. If there is no data from the requested categories,
	// then the results returned are empty and you're not charged.
	SourceCategoryRequirements SourceCategoryRequirements
}

// NewClient creates a new search client to submit queries with.
// Parameters values are set to the defaults defined by Pipl.
// For more information:
// https://docs.pipl.com/reference#configuration-parameters
func NewClient(APIKey string) (client *Client) {
	piplClient := new(Client)
	piplClient.HTTPClient = new(http.Client)
	piplClient.SearchParameters = new(SearchParameters)
	piplClient.SearchParameters.APIKey = APIKey
	piplClient.SearchParameters.MinimumProbability = 0.9
	piplClient.SearchParameters.InferPersons = false
	piplClient.SearchParameters.MinimumMatch = 0.0
	piplClient.SearchParameters.ShowSources = ShowSourcesNone
	piplClient.SearchParameters.HideSponsored = false
	piplClient.SearchParameters.LiveFeeds = true
	piplClient.SearchParameters.MatchRequirements = MatchRequirementsNone
	piplClient.SearchParameters.SourceCategoryRequirements = SourceCategoryRequirementsNone
	return piplClient
}

// meetsMinimumCriteria is used internally by SearchByPerson to do some very
// basic verification that the verify that search object has enough terms to
// meet the requirements for a search.
// From Pipl documentation:
// 		"The minimal requirement to run a search is to have at least one full
//		name, email, phone, username, user_id, URL or a single valid US address
//		(down to a house number). We can’t search for a job title or location
//		alone. We’re not a directory and can't provide bulk lists of people,
//		rather we specialize in identity resolution of single individuals."
func meetsMinimumCriteria(searchObject *Person) bool {
	if len(searchObject.Names) > 0 {
		for _, name := range searchObject.Names {
			if ((name.First != "") && (name.Last != "")) || (name.Raw != "") {
				return true
			}
		}
	}
	if len(searchObject.Emails) > 0 {
		for _, email := range searchObject.Emails {
			if email.Address != "" {
				return true
			}
		}
	}
	if len(searchObject.Phones) > 0 {
		for _, phone := range searchObject.Phones {
			if ((phone.CountryCode != 0) && (phone.Number != 0)) || (phone.Raw != "") {
				return true
			}
		}
	}
	if len(searchObject.Usernames) > 0 {
		for _, username := range searchObject.Usernames {
			if username.Content != "" {
				return true
			}
		}
	}
	if len(searchObject.UserIDs) > 0 {
		for _, userID := range searchObject.UserIDs {
			if userID.Content != "" {
				return true
			}
		}
	}
	if len(searchObject.URLs) > 0 {
		for _, url := range searchObject.URLs {
			if url.URL != "" {
				return true
			}
		}
	}
	return false
}

// SearchByPerson takes a person object (filled with search terms) and returns the
// results in the form of a Response struct. If successful, the response struct
// will contains the results, and err will be nil. If an error occurs, the struct pointer
// will be nil and you should check err for additional information.
func (searchClient *Client) SearchByPerson(searchObject *Person) (*Response, error) {
	if !meetsMinimumCriteria(searchObject) {
		return nil, &ErrInsufficientSearch{}
	}
	postData := url.Values{}
	postData.Add("key", searchClient.SearchParameters.APIKey)

	if searchClient.SearchParameters.ShowSources != ShowSourcesNone {
		postData.Add("show_sources", string(searchClient.SearchParameters.ShowSources))
	}
	if searchClient.SearchParameters.MatchRequirements != MatchRequirementsNone {
		postData.Add("match_requirements", string(searchClient.SearchParameters.MatchRequirements))
	}
	if searchClient.SearchParameters.SourceCategoryRequirements != SourceCategoryRequirementsNone {
		postData.Add("source_category_requirements", string(searchClient.SearchParameters.SourceCategoryRequirements))
	}
	personJSON, err := json.Marshal(searchObject)
	if err != nil {
		return nil, err
	}
	postData.Add("person", string(personJSON))
	request, err := http.NewRequest("POST", PiplAPIEndpoint, strings.NewReader(postData.Encode()))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := searchClient.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	piplResponse := new(Response)
	err = json.Unmarshal(body, piplResponse)
	if err != nil {
		return nil, err
	}
	return piplResponse, nil
}

// SearchByPointer takes a search pointer string and returns the full
// information for the person associated with that pointer
func (searchClient *Client) SearchByPointer(searchPointer string) (*Person, error) {
	postData := url.Values{}
	postData.Add("key", searchClient.SearchParameters.APIKey)
	postData.Add("search_pointer", searchPointer)
	request, err := http.NewRequest("POST", PiplAPIEndpoint, strings.NewReader(postData.Encode()))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := searchClient.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	piplResponse := new(Response)
	err = json.Unmarshal(body, piplResponse)
	if err != nil {
		return nil, err
	}
	return &piplResponse.Person, nil
}
