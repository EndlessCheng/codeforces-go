package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/4/C
// https://codeforces.com/problemset/status/4/problem/C
func TestCF4C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
abacaba
acaba
abacaba
acab
outputCopy
OK
OK
abacaba1
OK
inputCopy
6
first
first
second
second
third
third
outputCopy
OK
first1
OK
second1
OK
third1`
	testutil.AssertEqualCase(t, rawText, 0, CF4C)
}
