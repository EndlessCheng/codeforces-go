package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/965/D
// https://codeforces.com/problemset/status/965/problem/D
func TestCF965D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10 5
0 0 1 0 2 0 0 1 0
outputCopy
3
inputCopy
10 3
1 1 1 1 2 1 1 1 1
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF965D)
}
