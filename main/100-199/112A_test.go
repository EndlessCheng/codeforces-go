package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/112/A
// https://codeforces.com/problemset/status/112/problem/A
func TestCF112A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
aaaa
aaaA
outputCopy
0
inputCopy
abs
Abz
outputCopy
-1
inputCopy
abcdefg
AbCdEfF
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF112A)
}
