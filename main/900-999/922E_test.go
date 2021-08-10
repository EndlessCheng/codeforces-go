package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/922/E
// https://codeforces.com/problemset/status/922/problem/E
func TestCF922E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 12 0 4
3 4
4 2
outputCopy
6
inputCopy
4 1000 10 35
1 2 4 5
1000 500 250 200
outputCopy
5
inputCopy
2 10 7 11
2 10
6 1
outputCopy
11`
	testutil.AssertEqualCase(t, rawText, 0, CF922E)
}
