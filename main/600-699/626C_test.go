package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/626/C
// https://codeforces.com/problemset/status/626/problem/C
func TestCF626C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 3
outputCopy
9
inputCopy
3 2
outputCopy
8
inputCopy
5 0
outputCopy
10`
	testutil.AssertEqualCase(t, rawText, 0, CF626C)
}
