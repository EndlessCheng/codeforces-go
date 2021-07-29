package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/965/E
// https://codeforces.com/problemset/status/965/problem/E
func TestCF965E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
codeforces
codehorses
code
outputCopy
6
inputCopy
5
abba
abb
ab
aa
aacada
outputCopy
11
inputCopy
3
telegram
digital
resistance
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF965E)
}
