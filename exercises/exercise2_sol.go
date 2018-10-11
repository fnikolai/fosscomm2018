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
	checkForLanguage string = "python2"
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

func (g gentleman) Name() string {
	return g.name
}

func (g gentleman) KnowsPerl() bool {
	return g.lingua == "perl"
}

func (g gentleman) KnowsPython() bool {
	return g.lingua == "python"
}

func (g gentleman) KnowsC() bool {
	return g.lingua == "c"
}

type Attendee interface {
	Name() string
	KnowsPerl() bool
	KnowsPython() bool
	KnowsC() bool
}

// Lady ... gentleman ... does not matter what is the struct as long
// as it implements the specific set of methods
func scan(attendees map[string]Attendee, language string) (knowers []string, err error) {
	if attendees == nil {
		return nil, errors.New("NIL map was given")
	}

	defer func() {
		msg := recover()
		if msg != nil {
			fmt.Println("Recover from panic errror:", msg)
			knowers = nil
			err = errors.New(msg.(string))
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

	// Evaluate arguments on the entrace, calls function on the exit
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
