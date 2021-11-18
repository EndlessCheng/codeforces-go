package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/525/B
// https://codeforces.com/problemset/status/525/problem/B
func TestCF525B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
abcdef
1
2
outputCopy
aedcbf
inputCopy
vwxyz
2
2 2
outputCopy
vwxyz
inputCopy
abcdef
3
1 2 3
outputCopy
fbdcea`
	testutil.AssertEqualCase(t, rawText, 0, CF525B)
}
