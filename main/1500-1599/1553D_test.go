package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1553/D
// https://codeforces.com/problemset/status/1553/problem/D
func TestCF1553D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
ababa
ba
ababa
bb
aaa
aaaa
aababa
ababa
outputCopy
YES
NO
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1553D)
}
