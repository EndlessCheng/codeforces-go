package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1110/problem/D
// https://codeforces.com/problemset/status/1110/problem/D
func TestCF1110D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10 6
2 3 3 3 4 4 4 5 5 6
outputCopy
3
inputCopy
12 6
1 5 3 3 3 4 3 5 3 2 3 3
outputCopy
3
inputCopy
13 5
1 1 5 1 2 3 3 2 4 2 3 4 5
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1110D)
}
