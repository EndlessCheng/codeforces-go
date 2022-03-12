package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1650/problem/A
// https://codeforces.com/problemset/status/1650/problem/A
func TestCF1650A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
abcde
c
abcde
b
x
y
aaaaaaaaaaaaaaa
a
contest
t
outputCopy
YES
NO
NO
YES
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1650A)
}
