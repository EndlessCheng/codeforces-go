package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/677/D
// https://codeforces.com/problemset/status/677/problem/D
func TestCF677D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 4 3
2 1 1 1
1 1 1 1
2 1 1 3
outputCopy
5
inputCopy
3 3 9
1 3 5
8 9 7
4 6 2
outputCopy
22
inputCopy
3 4 12
1 2 3 4
8 7 6 5
9 10 11 12
outputCopy
11`
	testutil.AssertEqualCase(t, rawText, 0, CF677D)
}
