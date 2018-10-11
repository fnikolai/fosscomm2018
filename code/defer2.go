package main

func main() {
	f, err := os.Create("file")
	if err != nil {
		return
		panic("cannot create file")
	}
	defer f.Close()
	// no matter what happens here file will be closed
	// for sake of simplicity I skip checking close result
	fmt.Fprintf(f, "hello")
}
