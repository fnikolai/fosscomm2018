package main

import (
	"errors"
	"fmt"
)

var (
	attendees        map[string]Attendee
	knownLanguages          = []string{"perl", "python", "c"}
	checkForLanguage string = "python"
)


type Person struct {
	Name string
	Age int	
}

type Attendee struct {
	/* FILL: Struct can be built as composition of other structs */
	language string
}

// And structs can bind to methods (no inheritance though)
func (a Attendee) KnowsPerl() bool {
	return a.language=="perl"
}

func (a Attendee) KnowsPython() bool {
	return a.language=="python"
}

func (a Attendee) KnowsC() bool {
	return a.language=="c"
}



// Scan: finds the attendees who know the language and returns a list with their names
func scan(attendees map[string]Attendee, language string) ([]string, error) {
	// error is a value
	if attendees == nil {
		return nil, errors.New("NIL map was given")
	}

	if !stringInSlice(language, knownLanguages) {
		return nil, errors.New("invalid language")
	}


	// Declare a slice 
	var knowers []string

	// for map, range returns a tuple <key, value>
	for _, attendee := range attendees {
	
		// Simpler than if-else-if-else
		switch {
			case language == "perl" :
				if attendee.KnowsPerl() {
					/* FILL: append() automatically allocates the needed space for slice */
				}
				
			case language == "python" :
			 	if attendee.KnowsPython() {
					/* FILL: append() automatically allocates the needed space for slice */
				}
				
			case language == "c" :
				if attendee.KnowsC() {
					/* FILL: append() automatically allocates the needed space for slice */
				}
			default:
				return nil, errors.New("Invalid language")	
		}
	}
	
	if len(knowers) == 0 {
		return nil, errors.New("no one knows this language, sorry")
	}

	/* FILL: Return the knower list and denote successe */
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
	if /* FILL: error checking */ {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Ret:", ret)
}


func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
