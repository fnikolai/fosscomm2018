package main

import (
	"fmt"
	"runtime"
)

func main() {
	// START OMIT
	switch os := runtime.GOOS; os { // HL
	case "darwin":
		fmt.Println("OS X.")

	case "linux":
		fmt.Println("Linux.")

	case "*bsd", "plan9", "windows": // HL
		fmt.Println("Known others")

	default: // HL
		fmt.Printf("Unknown other %s.", os)
	}
	// END OMIT
}
