package main

import "fmt"

type Person struct {
	Name  string
	Actor string
}

func main() {
	// START OMIT
	var archenemy = struct {
		Person
		Nationality string
	}{Person{"Austin Powers", "Mike Myers"}, "British"}
	fmt.Println(archenemy)
	// END OMIT
}
