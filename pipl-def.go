// Package pipl provides a way to interact programatically with the Pipl API in Golang.
// For more detailed information on the Pipl search API and what we're actually
// wrapping, check out their official API reference: https://docs.pipl.com/reference/#overview
package pipl

// GUID is a unique format (but is just a string internally, since there's currently
// nothing all that fancy done with GUIDs). Additional guid-handling code may be
// added at a later date if needed.
type GUID string

// Name fields collectively define a possible name for a given person.
// If a search did not return information for a given field, it will be empty.
type Name struct {
	ValidSince string `json:"@valid_since,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	Current    bool   `json:"@current,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
	Type       string `json:"@type,omitempty"`
	First      string `json:"first,omitempty"`
	Middle     string `json:"middle,omitempty"`
	Last       string `json:"last,omitempty"`
	Prefix     string `json:"prefix,omitempty"`
	Suffix     string `json:"suffix,omitempty"`
	Raw        string `json:"raw,omitempty"`
	Display    string `json:"display,omitempty"`
}

// Address fields collectively define a possible address for a given person
// If a search did not return information for a given field, it will be empty.
type Address struct {
	ValidSince string `json:"@valid_since,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	Current    bool   `json:"@current,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
	Type       string `json:"@type,omitempty"`
	Country    string `json:"country,omitempty"`
	State      string `json:"state,omitempty"`
	City       string `json:"city,omitempty"`
	Street     string `json:"street,omitempty"`
	House      string `json:"house,omitempty"`
	Apartment  string `json:"apartment,omitempty"`
	ZipCode    string `json:"zip_code,omitempty"`
	POBox      string `json:"po_box,omitempty"`
	Raw        string `json:"raw,omitempty"`
	Display    string `json:"display,omitempty"`
}

// Phone fields collectively define a possible phone number for a given person
// If a search did not return information for a given field, it will be empty.
type Phone struct {
	ValidSince           string `json:"@valid_since,omitempty"`
	LastSeen             string `json:"@last_seen,omitempty"`
	Current              bool   `json:"@current,omitempty"`
	Inferred             bool   `json:"@inferred,omitempty"`
	Type                 string `json:"@type,omitempty"`
	CountryCode          int    `json:"country_code,omitempty"`
	Number               int    `json:"number,omitempty"`
	Extension            int    `json:"extension,omitempty"`
	Raw                  string `json:"raw,omitempty"`
	Display              string `json:"display,omitempty"`
	DisplayInternational string `json:"display_international,omitempty"`
}

// Email fields collectively define a possible email address for a given person
// If a search did not return information for a given field, it will be empty.
type Email struct {
	ValidSince    string `json:"@valid_since,omitempty"`
	LastSeen      string `json:"@last_seen,omitempty"`
	Current       bool   `json:"@current,omitempty"`
	Inferred      bool   `json:"@inferred,omitempty"`
	Type          string `json:"@type,omitempty"`
	Address       string `json:"address,omitempty"`
	AddressMD5    string `json:"address_md5,omitempty"`
	Disposable    bool   `json:"@disposable,omitempty"`
	EmailProvider bool   `json:"@email_provider,omitempty"`
}

// Username fields collectively define a possible username used by a given person.
// If a search did not return information for a given field, it will be empty.
type Username struct {
	ValidSince string `json:"@valid_since,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	Current    bool   `json:"@current,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
	Content    string `json:"content,omitempty"`
}

// UserID fields collectively define a possible UserID used by a given person.
// If a search did not return information for a given field, it will be empty.
type UserID struct {
	ValidSince string `json:"@valid_since,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	Current    bool   `json:"@current,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
	Content    string `json:"content,omitempty"`
}

// DateRange specifies a range of time by a start and end date
type DateRange struct {
	ValidSince string `json:"@valid_since,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	Current    bool   `json:"@current,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
	Start      string `json:"start,omitempty"`
	End        string `json:"end,omitempty"`
}

// DateOfBirth specififes a possible DOB for a person.
type DateOfBirth struct {
	ValidSince string    `json:"@valid_since,omitempty"`
	LastSeen   string    `json:"@last_seen,omitempty"`
	Current    bool      `json:"@current,omitempty"`
	Inferred   bool      `json:"@inferred,omitempty"`
	DateRange  DateRange `json:"date_range,omitempty"`
	Display    string    `json:"display,omitempty"`
}

// Image specifies a link to an image closely associated with the given person.
type Image struct {
	ValidSince     string `json:"@valid_since,omitempty"`
	LastSeen       string `json:"@last_seen,omitempty"`
	Current        bool   `json:"@current,omitempty"`
	Inferred       bool   `json:"@inferred,omitempty"`
	URL            string `json:"url,omitempty"`
	ThumbnailToken string `json:"thumbnail_token,omitempty"`
}

// Job specifies information about a possible occupation held by the given person.
type Job struct {
	ValidSince   string    `json:"@valid_since,omitempty"`
	LastSeen     string    `json:"@last_seen,omitempty"`
	Current      bool      `json:"@current,omitempty"`
	Inferred     bool      `json:"@inferred,omitempty"`
	Title        string    `json:"title,omitempty"`
	Organization string    `json:"organization,omitempty"`
	Industry     string    `json:"industry,omitempty"`
	DateRange    DateRange `json:"date_range,omitempty"`
	Display      string    `json:"display,omitempty"`
}

// Education specifies a possible
type Education struct {
	ValidSince string    `json:"@valid_since,omitempty"`
	LastSeen   string    `json:"@last_seen,omitempty"`
	Current    bool      `json:"@current,omitempty"`
	Inferred   bool      `json:"@inferred,omitempty"`
	Degree     string    `json:"degree,omitempty"`
	School     string    `json:"school,omitempty"`
	DateRange  DateRange `json:"date_range,omitempty"`
	Display    string    `json:"display,omitempty"`
}

// Gender contains a  possible gender of the given person.
// Gender is one of: "male", "female" (There is no default value for this field)
type Gender struct {
	ValidSince string `json:"@valid_since,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	Current    bool   `json:"@current,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
	Content    string `json:"content,omitempty"`
}

// Ethnicity contains a possible ethnicity of given person.
type Ethnicity struct {
	ValidSince string `json:"@valid_since,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	Current    bool   `json:"@current,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
	Content    string `json:"content,omitempty"`
}

// Language contains information about a possible language known by the given person.
type Language struct {
	ValidSince string `json:"@valid_since,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	Current    bool   `json:"@current,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
	Language   string `json:"language,omitempty"`
	Region     string `json:"region,omitempty"`
	Display    string `json:"display,omitempty"`
}

// OriginCountry contains information about a possible origin country of the
// given person.
type OriginCountry struct {
	ValidSince string `json:"@valid_since,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	Current    bool   `json:"@current,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
	Country    string `json:"country,omitempty"`
}

// Relationship contains information about a person who is closely related to
// the person being searched. This can be family members, spouses, children, etc.
// Type  and Subtype contain information about the nature of the relationship to
// the person being searched. For example, Type = "Family", Subtype = "Father".
// Type can be one of: "work", "family", "friend" (default), "other"
type Relationship struct {
	ValidSince      string          `json:"@valid_since,omitempty"`
	LastSeen        string          `json:"@last_seen,omitempty"`
	Current         bool            `json:"@current,omitempty"`
	Inferred        bool            `json:"@inferred,omitempty"`
	Type            string          `json:"@type,omitempty"`
	Subtype         string          `json:"@subtype,omitempty"`
	Names           []Name          `json:"names,omitempty"`
	Emails          []Email         `json:"emails,omitempty"`
	Usernames       []Username      `json:"usernames,omitempty"`
	Phones          []Phone         `json:"phones,omitempty"`
	Gender          Gender          `json:"gender,omitempty"`
	DateOfBirth     DateOfBirth     `json:"dob,omitempty"`
	Languages       []Language      `json:"languages,omitempty"`
	Ethnicities     []Ethnicity     `json:"ethnicities,omitempty"`
	OriginCountries []OriginCountry `json:"origin_countries,omitempty"`
	Addresses       []Address       `json:"addresses,omitempty"`
	Jobs            []Job           `json:"jobs,omitempty"`
	Educations      []Education     `json:"educations,omitempty"`
	Relationships   []Relationship  `json:"relationships,omitempty"`
	UserIDs         []UserID        `json:"user_ids,omitempty"`
	Images          []Image         `json:"images,omitempty"`
}

// URL contains information about a URL that is closely associated with a given person.
type URL struct {
	ValidSince string `json:"@valid_since,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	Current    bool   `json:"@current,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
	SourceID   string `json:"@source_id,omitempty"`
	Domain     string `json:"@domain,omitempty"`
	Name       string `json:"@name,omitempty"`
	Category   string `json:"@category,omitempty"`
	URL        string `json:"url,omitempty"`
}

// Tag contains content classification information
type Tag struct {
	Classification string `json:"@classification,omitempty"`
	Content        string `json:"content,omitempty"`
}

// Person contains all the information pertaining to a possible person match,
// including potential multiples of basic fields (names, emails, jobs, etc).
// The Match field represents the confidence of a particular person match, as a
// float: 0 <= Match <= 1. More potential matches returned in a search decreases
// the overall confidence of all matches.
type Person struct {
	ID              GUID            `json:"@id,omitempty"`
	Inferred        bool            `json:"@inferred,omitempty"`
	SearchPointer   string          `json:"@search_pointer,omitempty"`
	Match           float32         `json:"@match,omitempty"`
	Names           []Name          `json:"names,omitempty"`
	Emails          []Email         `json:"emails,omitempty"`
	Usernames       []Username      `json:"usernames,omitempty"`
	Phones          []Phone         `json:"phones,omitempty"`
	Gender          *Gender         `json:"gender,omitempty"`
	DateOfBirth     *DateOfBirth    `json:"dob,omitempty"`
	Languages       []Language      `json:"languages,omitempty"`
	Ethnicities     []Ethnicity     `json:"ethnicities,omitempty"`
	OriginCountries []OriginCountry `json:"origin_countries,omitempty"`
	Addresses       []Address       `json:"addresses,omitempty"`
	Jobs            []Job           `json:"jobs,omitempty"`
	Educations      []Education     `json:"educations,omitempty"`
	Relationships   []Relationship  `json:"relationships,omitempty"`
	UserIDs         []UserID        `json:"user_ids,omitempty"`
	URLs            []URL           `json:"urls,omitempty"`
}

// Source contains all the information for a given person, gathered from a
// single source. The source structure contains information about the name,
// domain, category, and source URL (amongst other fields).
type Source struct {
	ID              string          `json:"@id"`
	Name            string          `json:"@name"`
	Category        string          `json:"@category"`
	Domain          string          `json:"@domain"`
	PersonID        GUID            `json:"@person_id"`
	Sponsored       bool            `json:"@sponsored"`
	OriginURL       string          `json:"@origin_url"`
	Match           float32         `json:"@match"`
	Premium         bool            `json:"@premium"`
	Names           []Name          `json:"names"`
	Emails          []Email         `json:"emails"`
	Usernames       []Username      `json:"usernames"`
	Phones          []Phone         `json:"phones"`
	Gender          Gender          `json:"gender"`
	DateOfBirth     DateOfBirth     `json:"dob"`
	Languages       []Language      `json:"languages"`
	Ethnicities     []Ethnicity     `json:"ethnicities"`
	OriginCountries []OriginCountry `json:"origin_countries"`
	Addresses       []Address       `json:"addresses"`
	Jobs            []Job           `json:"jobs"`
	Educations      []Education     `json:"educations"`
	Relationships   []Relationship  `json:"relationships"`
	UserIDs         []UserID        `json:"user_ids"`
	URLs            []URL           `json:"urls"`
	Tags            []Tag           `json:"tags"`
}

// FieldCount contains the count of various attributes returned from a search
type FieldCount struct {
	Emails          int `json:"emails"`
	Relationships   int `json:"relationships"`
	Usernames       int `json:"usernames"`
	UserIDs         int `json:"user_ids"`
	Jobs            int `json:"jobs"`
	Addresses       int `json:"addresses"`
	Ethnicities     int `json:"ethnicities"`
	Phones          int `json:"phones"`
	MobilePhones    int `json:"mobile_phones"`
	LandlinePhones  int `json:"landline_phones"`
	Educations      int `json:"educations"`
	Languages       int `json:"languages"`
	SocialProfiles  int `json:"social_profiles"`
	Names           int `json:"names"`
	DOBs            int `json:"dobs"`
	Images          int `json:"images"`
	Genders         int `json:"genders"`
	OriginCountries int `json:"origin_countries"`
}

// AvailableData aggregates the counts for found attributes that are relevant to
// your search, divided into free and paid sources.
type AvailableData struct {
	Basic   FieldCount `json:"basic"`
	Premium FieldCount `json:"premium"`
}

// Response holds search results and general request information returned from
// the Pipl API. If an error occurs, the Error field will have more information.
// A search may be successful, but have some warnings. These are held in the
// Warnings field.
type Response struct {
	HTTPStatusCode    int           `json:"@http_status_code"`
	VisibleSources    int           `json:"@visible_sources"`
	AvailableSources  int           `json:"@available_sources"`
	PersonsCount      int           `json:"@persons_count"`
	SearchID          string        `json:"@search_id"`
	Query             Person        `json:"query"`
	MatchRequirements string        `json:"match_requirements"`
	AvailableData     AvailableData `json:"available_data"`
	Error             string        `json:"error"`
	Warnings          []string      `json:"warnings"`
	Person            Person        `json:"person"`
	PossiblePersons   []Person      `json:"possible_persons"`
	Sources           []Source      `json:"sources"`
}
