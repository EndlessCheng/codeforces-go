package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/25/E
// https://codeforces.com/problemset/status/25/problem/E
func TestCF25E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
ab
bc
cd
outputCopy
4
inputCopy
abacaba
abaaba
x
outputCopy
11`
	testutil.AssertEqualCase(t, rawText, 0, CF25E)
}
