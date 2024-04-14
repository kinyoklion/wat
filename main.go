package main

import (
	"fmt"
	"slices"
	"wat/algorithms"
	wat "wat/internal"
)

func main() {
	distance := algorithms.LevenshteinDistance("sitting", "kitten")
	fmt.Println("Distance:", distance)

	distance = algorithms.LevenshteinDistance("Saturday", "Sunday")
	fmt.Println("Distance:", distance)

	a := wat.Entry{Details: "sitting"}
	b := wat.Entry{Details: "kitten"}

	items := []wat.Entry{a, b}
	slices.SortFunc(items, func(a, b wat.Entry) int {
		return algorithms.LevenshteinCmp("sitting", a.Details, b.Details)
	})

	fmt.Println(items)

}
