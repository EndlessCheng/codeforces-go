package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/633/B
// https://codeforces.com/problemset/status/633/problem/B
func TestCF633B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
outputCopy
5
5 6 7 8 9 
inputCopy
5
outputCopy
0
inputCopy
100000
outputCopy
`
	testutil.AssertEqualCase(t, rawText, 0, CF633B)
}
