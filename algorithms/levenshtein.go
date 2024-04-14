package algorithms

import (
	"fmt"
	"unicode/utf8"
)

// Implementation of the Wagner–Fischer algorithm
// Source: https://en.wikipedia.org/wiki/Wagner%E2%80%93Fischer_algorithm

func LevenshteinDistance(s, t string) int {
	m := utf8.RuneCountInString(s) + 1 + 1
	n := utf8.RuneCountInString(t) + 1 + 1

	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}

	for i := 1; i < m; i++ {
		matrix[i][0] = i
	}

	for j := 0; j < n; j++ {
		matrix[0][j] = j
	}

	for j := 0; j < utf8.RuneCountInString(t); j++ {
		for i := 0; i < utf8.RuneCountInString(s); i++ {
			subCost := 1
			if s[i] == t[j] {
				subCost = 0
			}

			matrixI := i + 1
			matrixJ := j + 1
			matrix[matrixI][matrixJ] = min(
				matrix[matrixI-1][matrixJ]+1,
				matrix[matrixI][matrixJ-1]+1,
				matrix[matrixI-1][matrixJ-1]+subCost,
			)
		}
	}

	fmt.Println(matrix)

	return matrix[len(s)][len(t)]
}

func LevenshteinCmp(ref string, a string, b string) int {
	return LevenshteinDistance(ref, a) - LevenshteinDistance(ref, b)
}
