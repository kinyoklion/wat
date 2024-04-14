package collections

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	doubled := Map(items, func(item int) int {
		return item * 2
	})

	expectedDoubled := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	if !reflect.DeepEqual(expectedDoubled, doubled) {
		t.Fatalf("got %v, expected %v", doubled, expectedDoubled)
	}

	strings := Map(items, func(item int) string {
		return strconv.Itoa(item)
	})

	expectedStrings := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	if !reflect.DeepEqual(expectedStrings, strings) {
		t.Fatalf("got %v, expected %v", strings, expectedStrings)
	}
}
