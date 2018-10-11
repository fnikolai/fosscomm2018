// Exercise 3 demonstrates:
//
// 1) goroutines
// 2) goroutines synchronization
// 3) asynchronous channels

package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	attendees        map[string]Attendee
	knownLanguages          = []string{"perl", "python", "c"}
	checkForLanguage string = "python"
)

func init() {
	attendees = map[string]Attendee{
		"l0": lady{"Maria", knownLanguages[0], nil},
		"l1": lady{name: "Eleni", language: knownLanguages[1]},
		"l2": lady{name: "Katerina", language: knownLanguages[2], information: []struct{}{}},

		"g0": gentleman{"Fotis", knownLanguages[0]},
		"g1": gentleman{name: "Manos", lingua: knownLanguages[1]},
		"g2": gentleman{lingua: knownLanguages[2], name: "Giorgos"},
	}
}

// Because the slice will be accessed by multiple goroutines
// it must be protected with a lock (struct composition)
type Knowers struct {
	sync.Mutex
	names []string
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

func scan(attendees map[string]Attendee, language string, retCh chan []string, errorCh chan error) {
	// Return errors through channels. In this case, we do not need to proceed
	// further within the function
	if attendees == nil {
		errorCh <- errors.New("NIL map was given")
		return
	}

	defer func() {
		msg := recover()
		if msg != nil {
			fmt.Println("Recover from panic errror:", msg)
			errorCh <- errors.New(msg.(string))
		}
	}()

	// declare a dynamic group for waiting goroutines to terminate
	// like a semaphore
	wg := sync.WaitGroup{}
	knowers := Knowers{}

	for _, attendee := range attendees {
		// For every goroutine incresae the semaphore
		// Decrease it when the goroutine terminates
		wg.Add(1)
		go func(attendee Attendee) {
			defer wg.Done()

			switch {
			case language == "perl":
				if attendee.KnowsPerl() {
					knowers.Lock()
					knowers.names = append(knowers.names, attendee.Name())
					knowers.Unlock()
				}

			case language == "python":
				if attendee.KnowsPython() {
					knowers.Lock()
					knowers.names = append(knowers.names, attendee.Name())
					knowers.Unlock()

				}

			case language == "c":
				if attendee.KnowsC() {
					knowers.Lock()
					knowers.names = append(knowers.names, attendee.Name())
					knowers.Unlock()
				}
			default:
				panic("invalid language")
			}
		}(attendee)
	}

	// Block until all of the spawned goroutines have finished
	wg.Wait()

	// return either an error or the result through channels
	if len(knowers.names) == 0 {
		errorCh <- errors.New("no one knows this language, sorry")
		return
	}
	retCh <- knowers.names
}

func main() {
	fmt.Println("Hello, playground")

	defer timeTrack(time.Now(), "Scanning")

	// Initiate asynchronous channels of particular type
	retCh, errorCh := make(chan []string), make(chan error)

	// spawn the function into a new goroutine and funnel returned
	// values through those channels
	go scan(attendees, checkForLanguage, retCh, errorCh)

	//listen the channels for returned values
	select {
	case knowers := <-retCh:
		fmt.Println("People who know this language:", knowers)
	case err := <-errorCh:
		fmt.Println("Error:", err)
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %v sec", name, elapsed.Seconds())
}
