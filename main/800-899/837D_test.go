package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/837/D
// https://codeforces.com/problemset/status/837/problem/D
func TestCF837D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
50 4 20
outputCopy
3
inputCopy
5 3
15 16 3 25 9
outputCopy
3
inputCopy
3 3
9 77 13
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF837D)
}
