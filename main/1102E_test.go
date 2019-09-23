package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1102E(t *testing.T) {
	// just copy from website
	rawText := `100
71 23 84 98 8 14 4 42 56 83 87 28 22 32 50 5 96 90 1 59 74 56 96 77 88 71 38 62 36 85 1 97 98 98 32 99 42 6 81 20 49 57 71 66 9 45 41 29 28 32 68 38 29 35 29 19 27 76 85 68 68 41 32 78 72 38 19 55 83 83 25 46 62 48 26 53 14 39 31 94 84 22 39 34 96 63 37 42 6 78 76 64 16 26 6 79 53 24 29 63
outputCopy
1
inputCopy
5
1 2 1 2 3
outputCopy
2
inputCopy
2
100 1
outputCopy
2
inputCopy
4
1 3 3 7
outputCopy
4`
	testutil.AssertEqual(t, rawText, Sol1102E)
}
