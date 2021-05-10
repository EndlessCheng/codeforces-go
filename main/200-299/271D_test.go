package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/271/D
// https://codeforces.com/problemset/status/271/problem/D
func TestCF271D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
ababab
01000000000000000000000000
1
outputCopy
5
inputCopy
acbacbacaa
00000000000000000000000000
2
outputCopy
8`
	testutil.AssertEqualCase(t, rawText, 0, CF271D)
}
