package collections

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	filtered := Filter(items, func(item int) bool {
		return item%2 == 0
	})

	if !reflect.DeepEqual([]int{2, 4, 6, 8, 10}, filtered) {
		t.Fatalf("Filter failed %v", filtered)
	}
}
