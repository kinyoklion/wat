package collections

import (
	"reflect"
	"testing"
)

func TestValues(t *testing.T) {
	sample := map[string]string{"foo": "bar", "bar": "baz"}

	expected := []string{"bar", "baz"}
	got := Values(sample)
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("Got %v, expected %v", got, expected)
	}
}
