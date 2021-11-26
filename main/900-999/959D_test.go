package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/959/D
// https://codeforces.com/problemset/status/959/problem/D
func TestCF959D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 3 5 4 13
outputCopy
2 3 5 7 11 
inputCopy
3
10 3 7
outputCopy
10 3 7 `
	testutil.AssertEqualCase(t, rawText, 0, CF959D)
}
