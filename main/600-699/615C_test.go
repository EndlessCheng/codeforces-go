package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/615/C
// https://codeforces.com/problemset/status/615/problem/C
func TestCF615C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
abc
cbaabc
outputCopy
2
3 1
1 3
inputCopy
aaabrytaaa
ayrat
outputCopy
3
1 1
6 5
8 7
inputCopy
ami
no
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF615C)
}
