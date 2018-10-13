package pipl_test

import (
	"fmt"
	"log"
	"os"

	"github.com/xpcmdshell/pipl"
)

var ()

func Example() {
	log.Println("Starting up.")
	// Create a new client object to handle your queries (supply an API Key)
	client := pipl.NewClient("API Key Goes Here")

	// Set your match requirements if you have any. You don't pay for results that
	// don't satisfy your match requirements (but your returned results will be empty)
	client.SearchParameters.MatchRequirements = "name and phone"

	// Create a blank person to fill out with search terms
	searchObject := pipl.NewPerson()

	// Let's find out who this random guy is. We'll search by a username.
	searchObject.AddUsername("@jeffbezos")

	// And we'll add a full name
	searchObject.AddName("jeff", "preston", "bezos", "", "")

	// Some field types have a "raw" option that you can use if you're unsure
	// how to break it down. Pipl will attempt to parse it for you.
	// Generally you should use one or the other (AddX() or AddXRaw())
	searchObject.AddNameRaw("jeff preston bezos")

	// Launch the search (if you don't meet the minimum search criteria, an error
	// should be returned to you here stating such).
	results, err := client.SearchByPerson(searchObject)
	// Handle errors better than I do pls
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// When multiple PossiblePersons are returned, we get a "preview" of each of
	// each of them (< 100% match confidence)
	if results.PersonsCount > 1 {
		for _, person := range results.PossiblePersons {
			// In order to get the full info on each, we need to a follow up query
			// to pull a full person profile by search pointer
			searchPtr := person.SearchPointer
			ptrResults, err := client.SearchByPointer(searchPtr)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			ptrSummary, err := ptrResults.Summarize()
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			fmt.Println(ptrSummary)
		}
	} else if results.PersonsCount == 1 {
		// When a single result is returned from our search, we get a full
		// profile by default (100% match confidence)
		personSummary, err := results.Person.Summarize()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		fmt.Println(personSummary)
	} else {
		fmt.Println("No results!")
	}
	log.Println("Shutting down.")
}
