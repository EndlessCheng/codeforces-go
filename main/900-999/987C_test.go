package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/987/C
// https://codeforces.com/problemset/status/987/problem/C
func TestCF987C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 4 5 4 10
40 30 20 10 40
outputCopy
90
inputCopy
3
100 101 100
2 4 5
outputCopy
-1
inputCopy
10
1 2 3 4 5 6 7 8 9 10
10 13 11 14 15 12 13 13 18 13
outputCopy
33`
	testutil.AssertEqualCase(t, rawText, 0, CF987C)
}
