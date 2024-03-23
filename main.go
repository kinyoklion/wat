package main

import (
	"log"
	internal "wat/internal"
)

func main() {
	form, _ := internal.EntryForm()

	err := form.Run()

	if err != nil {
		log.Fatal(err)
	}
}
