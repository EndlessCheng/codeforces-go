package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/182/D
// https://codeforces.com/problemset/status/182/problem/D
func TestCF182D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
abcdabcd
abcdabcdabcdabcd
outputCopy
2
inputCopy
aaa
aa
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF182D)
}
