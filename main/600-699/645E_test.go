package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/645/E
// https://codeforces.com/problemset/status/645/problem/E
func TestCF645E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 3
ac
outputCopy
8
inputCopy
0 2
aaba
outputCopy
10
inputCopy
5 3
aabcc
outputCopy
`
	testutil.AssertEqualCase(t, rawText, 0, CF645E)
}
