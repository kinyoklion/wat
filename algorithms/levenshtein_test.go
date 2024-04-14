package algorithms

import "testing"

func TestLevenshteinDistance(t *testing.T) {
	if LevenshteinDistance("sitting", "kitten") != 3 {
		t.Fail()
	}

	if LevenshteinDistance("kitten", "sitting") != 3 {
		t.Fail()
	}

	if LevenshteinDistance("Ryan", "Bryan") != 2 {
		t.Fail()
	}

	if LevenshteinDistance("Ryan", "ryan") != 1 {
		t.Fail()
	}

	if LevenshteinDistance("four", "") != 4 {
		t.Fail()
	}

	if LevenshteinDistance("", "four") != 4 {
		t.Fail()
	}

	if LevenshteinDistance("", "") != 0 {
		t.Fail()
	}

	if LevenshteinDistance("zero", "zero") != 0 {
		t.Fail()
	}
}

func TestLevenshteinCmp(t *testing.T) {
	if LevenshteinCmp("kitten", "kitten", "sitting") > 0 {
		t.Fail()
	}
	if LevenshteinCmp("kitten", "sitting", "kitten") < 0 {
		t.Fail()
	}

	if LevenshteinCmp("kitten", "kitten", "kitten") != 0 {
		t.Fail()
	}
}
