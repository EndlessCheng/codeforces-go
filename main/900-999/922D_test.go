package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/922/D
// https://codeforces.com/problemset/status/922/problem/D
func TestCF922D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
ssh
hs
s
hhhs
outputCopy
18
inputCopy
2
h
s
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF922D)
}
