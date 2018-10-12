// Exercise 2 demonstrates:
//
// 1) init function
// 2) empty struct
// 3) interfaces
// 4) deferred execution for time measurement
// 5) panic-recovery

package main

import (
	"errors"
	"fmt"
	"time"
)

var (
	attendees        map[string]Attendee
	knownLanguages          = []string{"perl", "python", "c"}
	checkForLanguage string = "python"
)

// Init() is called before main
func init() {
	// Many different ways to declare a struct
	// - by order (all fields must be explicitly included)
	// - by field (fields may be ommitted, order does not matter)
	attendees = map[string]Attendee{
		"l0": lady{"Maria", knownLanguages[0], nil},
		"l1": lady{name: "Eleni", language: knownLanguages[1]},
		"l2": lady{name: "Katerina", language: knownLanguages[2], information: []struct{}{}},

		"g0": gentleman{"Fotis", knownLanguages[0]},
		"g1": gentleman{name: "Manos", lingua: knownLanguages[1]},
		"g2": gentleman{lingua: knownLanguages[2], name: "Giorgos"},
	}
}

type lady struct {
	name        string
	language    string
	information []struct{}
}

func (l lady) Name() string {
	return l.name
}

func (l lady) KnowsPerl() bool {
	return l.language == "perl"
}

func (l lady) KnowsPython() bool {
	return l.language == "python"
}

func (l lady) KnowsC() bool {
	return l.language == "c"
}

type gentleman struct {
	name   string
	lingua string
}

/*
* FILL: For a gentleman to be an Attendee, he must
* fulfill certain criteria. What are those criteria ?
 */
type Attendee interface {
	Name() string
	KnowsPerl() bool
	KnowsPython() bool
	KnowsC() bool
}

func scan(attendees map[string]Attendee, language string) ([]string, error) {
	if attendees == nil {
		return nil, errors.New("NIL map was given")
	}

	// Declare a slice
	var knowers []string

	defer func() {
		msg := recover()
		if msg != nil {
			fmt.Println("Recover from panic errror:", msg)

			/*
			* FILL: defer postpones the execution of the anonymous function until
			* the surrounding function returns
			* ... What should be the returned of scan function ?
			* ... What kind of housekeeping cleanup should be implemented here ?
			 */
		}
	}()

	for _, attendee := range attendees {

		switch {
		case language == "perl":
			if attendee.KnowsPerl() {
				knowers = append(knowers, attendee.Name())
			}

		case language == "python":
			if attendee.KnowsPython() {
				knowers = append(knowers, attendee.Name())
			}

		case language == "c":
			if attendee.KnowsC() {
				knowers = append(knowers, attendee.Name())
			}
		default:
			panic("invalid language")
		}
	}

	if len(knowers) == 0 {
		return nil, errors.New("no one knows this language, sorry")
	}

	return knowers, err
}

func main() {
	fmt.Println("Hello, playground")

	// Just a candy ... a fancy way to measure performance
	// defer evaluates the arguments on entrace, but calls
	// the function when main exits.
	defer timeTrack(time.Now(), "Scanning")

	ret, err := scan(attendees, checkForLanguage)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("People who know this language:", ret)
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %v sec", name, elapsed.Seconds())
}
