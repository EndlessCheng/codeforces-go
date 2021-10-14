package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1553/B
// https://codeforces.com/problemset/status/1553/problem/B
func TestCF1553B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
abcdef
cdedcb
aaa
aaaaa
aab
baaa
ab
b
abcdef
abcdef
ba
baa
outputCopy
YES
YES
NO
YES
YES
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1553B)
}
