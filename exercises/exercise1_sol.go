// Exercise 1 demonstrates:
//
// 1) global variables
// 2) Struct composition
// 3) Struct methods
// 4) Function calling
// 5) maps and map iterators
// 6) lists and append
// 7) switch statements
// 8) error handling
package main

import (
	"errors"
	"fmt"
)

var (
	knownLanguages          = []string{"perl", "python", "c"}
	checkForLanguage string = "python"
)

type Person struct {
	Name string
	Age  int
}

// Struct can be built as composition of other structs
type Attendee struct {
	Person
	language string
}

// And structs can bind to methods (no inheritance though)
func (a Attendee) KnowsPerl() bool {
	return a.language == "perl"
}

func (a Attendee) KnowsPython() bool {
	return a.language == "python"
}

func (a Attendee) KnowsC() bool {
	return a.language == "c"
}

func scan(attendees map[string]Attendee, language string) ([]string, error) {
	// error is a value
	if attendees == nil {
		return nil, errors.New("NIL map was given")
	}

	// Declare a slice
	var knowers []string

	// for map, range returns a tuple <key, value>
	for _, attendee := range attendees {

		// Simpler than if-else-if-else
		switch {
		case language == "perl":
			if attendee.KnowsPerl() {
				// Append automatically allocates the needed space for slice
				knowers = append(knowers, attendee.Name)
			}

		case language == "python":
			if attendee.KnowsPython() {
				knowers = append(knowers, attendee.Name)
			}

		case language == "c":
			if attendee.KnowsC() {
				knowers = append(knowers, attendee.Name)
			}
		default:
			return nil, errors.New("invalid language")
		}
	}

	if len(knowers) == 0 {
		return nil, errors.New("no one knows this language, sorry")
	}

	return knowers, nil
}

func main() {
	fmt.Println("Hello, playground")

	attendees := map[string]Attendee{
		// Define values in field order - similar to function calling
		"at0": Attendee{Person{Name: "Maria", Age: 17}, knownLanguages[0]},

		// Define values with named fields - order does not matter
		"at1": Attendee{Person: Person{Name: "Eleni", Age: 56}, language: knownLanguages[1]},

		// Define values with named fields - order does not matter
		"at2": Attendee{language: knownLanguages[2], Person: Person{Name: "Katerina", Age: 21}},
	}

	// Always do error checking after the function calling
	ret, err := scan(attendees, checkForLanguage)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("People who know this language:", ret)
}
