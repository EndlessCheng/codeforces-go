package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/103/B
// https://codeforces.com/problemset/status/103/problem/B
func TestCF103B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 6
6 3
6 4
5 1
2 5
1 4
5 4
outputCopy
FHTAGN!
inputCopy
6 5
5 6
4 6
3 1
5 1
1 2
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF103B)
}
