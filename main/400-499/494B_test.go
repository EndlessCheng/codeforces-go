package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/494/B
// https://codeforces.com/problemset/status/494/problem/B
func TestCF494B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
ababa
aba
outputCopy
5
inputCopy
welcometoroundtwohundredandeightytwo
d
outputCopy
274201
inputCopy
ddd
d
outputCopy
12`
	testutil.AssertEqualCase(t, rawText, -1, CF494B)
}
