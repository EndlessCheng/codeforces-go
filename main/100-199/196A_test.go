package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/196/A
// https://codeforces.com/problemset/status/196/problem/A
func TestCF196A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
ababba
outputCopy
bbba
inputCopy
abbcbccacbbcbaaba
outputCopy
cccccbba`
	testutil.AssertEqualCase(t, rawText, 0, CF196A)
}
