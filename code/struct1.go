package main

import "fmt"

type Person struct {
	Name  string
	Actor string
}

func main() {
	// START OMIT
	var me Person
	me.Name = "Douglas Powers"
	me.Actor = "Mike Myers"

	minime := Person{
		Name:  "Doublas Powers Jr.",
		Actor: "Mike Myers",
	}

	var archenemy = struct {
		Person
		Nationality string
	}{Person{"Austin Powers", "Mike Myers"}, "British"}

	fmt.Printf("%v\n", me)
	fmt.Printf("%+v\n", minime)
	fmt.Printf("%#v\n", archenemy)
	// END OMIT
}
