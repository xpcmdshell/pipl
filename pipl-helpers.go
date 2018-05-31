package pipl

import (
	"fmt"
	"strings"
)

// ErrInsufficientSearch is an error type that may be returned by
// SearchByPerson which denotes that the search object provided does not meet
// the minimum requirements.
type ErrInsufficientSearch struct{}

func (err *ErrInsufficientSearch) Error() string {
	return "The search object submitted does not contain sufficient terms. Must have a complete entry for one of the following: Name, email, phone, username, userID, url"
}

// Summarize returns a string summary of the attributes of a person object
func (person Person) Summarize() (string, error) {
	builder := strings.Builder{}
	_, err := builder.WriteString(fmt.Sprintf("Match Confidence: %.f%%\n", person.Match*100))
	if err != nil {
		return "", err
	}
	if len(person.Names) > 0 {
		_, err = builder.WriteString("Names:\n")
		if err != nil {
			return "", err
		}
		for _, name := range person.Names {
			_, err = builder.WriteString(fmt.Sprintf("\t%s %s %s\n", name.First, name.Middle, name.Last))
			if err != nil {
				return "", err
			}
		}
	}

	if len(person.Emails) > 0 {
		_, err = builder.WriteString("Email Addresses:\n")
		if err != nil {
			return "", err
		}
		for _, email := range person.Emails {
			_, err = builder.WriteString(fmt.Sprintf("\t%s\n", email.Address))
			if err != nil {
				return "", err
			}
		}
	}

	if len(person.Usernames) > 0 {
		_, err = builder.WriteString("Usernames:\n")
		for _, username := range person.Usernames {
			_, err = builder.WriteString(fmt.Sprintf("\t%s\n", username.Content))
			if err != nil {
				return "", err
			}
		}
	}

	if len(person.Phones) > 0 {
		_, err = builder.WriteString("Phone Numbers:\n")
		if err != nil {
			return "", err
		}
		for _, phone := range person.Phones {
			_, err = builder.WriteString(fmt.Sprintf("\t%s\n", phone.Display))
			if err != nil {
				return "", err
			}
		}
	}

	if person.Gender != nil {
		_, err = builder.WriteString(fmt.Sprintf("Gender:\n\t%s\n", person.Gender.Content))
		if err != nil {
			return "", err
		}
	}

	if person.DateOfBirth != nil {
		_, err = builder.WriteString(fmt.Sprintf("Date of Birth:\n\t%s\n", person.DateOfBirth.Display))
		if err != nil {
			return "", err
		}
	}

	if len(person.Languages) > 0 {
		_, err = builder.WriteString("Languages:\n")
		if err != nil {
			return "", err
		}
		for _, language := range person.Languages {
			_, err = builder.WriteString(fmt.Sprintf("\t%s\n", language.Display))
			if err != nil {
				return "", err
			}
		}
	}

	if len(person.Ethnicities) > 0 {
		_, err = builder.WriteString("Ethnicities:\n")
		if err != nil {
			return "", err
		}
		for _, ethnicity := range person.Ethnicities {
			_, err = builder.WriteString(fmt.Sprintf("\t%s\n", ethnicity.Content))
			if err != nil {
				return "", err
			}
		}
	}

	if len(person.OriginCountries) > 0 {
		_, err = builder.WriteString("Origin Countries:\n")
		if err != nil {
			return "", err
		}
		for _, originCountry := range person.OriginCountries {
			_, err = builder.WriteString(fmt.Sprintf("\t%s\n", originCountry.Country))
			if err != nil {
				return "", err
			}
		}
	}

	if len(person.Addresses) > 0 {
		_, err = builder.WriteString("Addresses:\n")
		if err != nil {
			return "", err
		}
		for _, address := range person.Addresses {
			_, err = builder.WriteString(fmt.Sprintf("\t%s\n", address.Display))
			if err != nil {
				return "", err
			}
		}
	}

	if len(person.Jobs) > 0 {
		_, err = builder.WriteString("Jobs:\n")
		if err != nil {
			return "", err
		}
		for _, job := range person.Jobs {
			_, err = builder.WriteString(fmt.Sprintf("\t%s\n", job.Display))
			if err != nil {
				return "", err
			}
		}
	}

	if len(person.Educations) > 0 {
		_, err = builder.WriteString("Education:\n")
		if err != nil {
			return "", err
		}
		for _, education := range person.Educations {
			_, err = builder.WriteString(fmt.Sprintf("\t%s\n", education.Display))
			if err != nil {
				return "", err
			}
		}
	}

	if len(person.Relationships) > 0 {
		_, err = builder.WriteString("Relationships:\n")
		if err != nil {
			return "", err
		}
		for _, relation := range person.Relationships {
			_, err = builder.WriteString(fmt.Sprintf("\t%s (%s, %s)\n", relation.Names[0].Display, relation.Type, relation.Subtype))
			if err != nil {
				return "", err
			}
		}
	}

	if len(person.UserIDs) > 0 {
		_, err = builder.WriteString("User IDs:\n")
		for _, id := range person.UserIDs {
			_, err = builder.WriteString(fmt.Sprintf("\t%s\n", id.Content))
			if err != nil {
				return "", err
			}
		}
	}

	if len(person.URLs) > 0 {
		_, err = builder.WriteString("Related URLs:\n")
		if err != nil {
			return "", err
		}
		for _, url := range person.URLs {
			_, err = builder.WriteString(fmt.Sprintf("\t%s\n", url.URL))
			if err != nil {
				return "", err
			}
		}
	}
	return builder.String(), nil
}

// NewPerson makes a new blank person object to be filled with terms
func NewPerson() *Person {
	return new(Person)
}

// AddName adds a name to the search object. For well defined names. Omit unused fields.
func (searchObject *Person) AddName(firstName string, middleName string, lastName string, prefix string, suffix string) {
	newName := new(Name)
	newName.First = firstName
	newName.Middle = middleName
	newName.Last = lastName
	newName.Prefix = prefix
	newName.Suffix = suffix
	searchObject.Names = append(searchObject.Names, *newName)
}

// AddNameRaw can be used when you're unsure how to handle breaking down the name in
// question into its constituent parts. Basically, let Pipl handle parsing it.
func (searchObject *Person) AddNameRaw(fullName string) {
	newName := new(Name)
	newName.Raw = fullName
	searchObject.Names = append(searchObject.Names, *newName)
}

// AddEmail appends an email address to the specified search object
func (searchObject *Person) AddEmail(emailAddress string) {
	newEmail := new(Email)
	newEmail.Address = emailAddress
	searchObject.Emails = append(searchObject.Emails, *newEmail)
}

// AddUsername appends a username to the specified search object
func (searchObject *Person) AddUsername(username string) {
	newUsername := new(Username)
	newUsername.Content = username
	searchObject.Usernames = append(searchObject.Usernames, *newUsername)
}

// AddPhone appends a phone to the specified search object
func (searchObject *Person) AddPhone(phoneNumber int) {
	newPhone := new(Phone) // who dis
	newPhone.Number = phoneNumber
	searchObject.Phones = append(searchObject.Phones, *newPhone)
}

// SetGender sets the gender of the specified search object
func (searchObject *Person) SetGender(gender string) {
	newGender := new(Gender)
	newGender.Content = gender
	searchObject.Gender = newGender
}

// SetDateOfBirth sets the DOB of the specified search object
// DOB string format: "YYYY-MM-DD"
func (searchObject *Person) SetDateOfBirth(dob string) {
	newDOB := new(DateOfBirth)
	newDOB.DateRange.Start = dob
	newDOB.DateRange.End = dob
	searchObject.DateOfBirth = newDOB
}

// AddLanguage appends a language to the specified search object.
// Language is a 2 character language code (e.g. "en")
// Region  is a country code (e.g "US")
func (searchObject *Person) AddLanguage(languageCode string, regionCode string) {
	newLanguage := new(Language)
	newLanguage.Language = languageCode
	newLanguage.Region = regionCode
	searchObject.Languages = append(searchObject.Languages, *newLanguage)
}

// AddEthnicity appends an ethnicity to the specified search object
func (searchObject *Person) AddEthnicity(ethnicity string) {
	newEthnicity := new(Ethnicity)
	newEthnicity.Content = ethnicity
	searchObject.Ethnicities = append(searchObject.Ethnicities, *newEthnicity)
}

// AddOriginCountry appends an origin country to the specified search object
func (searchObject *Person) AddOriginCountry(countryCode string) {
	newCountry := new(OriginCountry)
	newCountry.Country = countryCode
	searchObject.OriginCountries = append(searchObject.OriginCountries, *newCountry)
}

// AddAddress appends an address to the specified search object
func (searchObject *Person) AddAddress(house string, street string, apartment string, city string, state string, country string, poBox string) {
	newAddress := new(Address)
	newAddress.House = house
	newAddress.Street = street
	newAddress.Apartment = apartment
	newAddress.City = city
	newAddress.State = state
	newAddress.Country = country
	newAddress.POBox = poBox
	searchObject.Addresses = append(searchObject.Addresses, *newAddress)
}

// AddAddressRaw can be used when many of the address parts are missing, or
// you're unsure how to split it up. Let Pipl handle parsing.
func (searchObject *Person) AddAddressRaw(fullAddress string) {
	newAddress := new(Address)
	newAddress.Raw = fullAddress
	searchObject.Addresses = append(searchObject.Addresses, *newAddress)
}

// AddJob appends a job entry to the specified search object
func (searchObject *Person) AddJob(title string, organization string, industry string, dateRangeStart string, dateRangeEnd string) {
	newJob := new(Job)
	newJob.Title = title
	newJob.Organization = organization
	newJob.Industry = industry
	newJob.DateRange.Start = dateRangeStart
	newJob.DateRange.End = dateRangeEnd
	searchObject.Jobs = append(searchObject.Jobs, *newJob)
}

// AddEducation appends an education entry to the specified search object
func (searchObject *Person) AddEducation(degree string, school string, dateRangeStart string, dateRangeEnd string) {
	newEducation := new(Education)
	newEducation.Degree = degree
	newEducation.School = school
	newEducation.DateRange.Start = dateRangeStart
	newEducation.DateRange.End = dateRangeEnd
	searchObject.Educations = append(searchObject.Educations, *newEducation)
}

// AddUserID appends a user ID to the specified search object
func (searchObject *Person) AddUserID(userID string) {
	newUserID := new(UserID)
	newUserID.Content = userID
	searchObject.UserIDs = append(searchObject.UserIDs, *newUserID)
}

// AddURL appends a URL to the specified search object
func (searchObject *Person) AddURL(url string) {
	newURL := new(URL)
	newURL.URL = url
	searchObject.URLs = append(searchObject.URLs, *newURL)
}

// AddRelationship appends a relationship entry to the specified search object
func (searchObject *Person) AddRelationship(relationship Relationship) {
	searchObject.Relationships = append(searchObject.Relationships, relationship)
}