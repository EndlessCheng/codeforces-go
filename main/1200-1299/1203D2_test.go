package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1203/D2
// https://codeforces.com/problemset/status/1203/problem/D2
func TestCF1203D2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
bbaba
bb
outputCopy
3
inputCopy
baaba
ab
outputCopy
2
inputCopy
abcde
abcde
outputCopy
0
inputCopy
asdfasdf
fasd
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1203D2)
}
