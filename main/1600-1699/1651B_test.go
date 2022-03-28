package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1651/problem/B
// https://codeforces.com/problemset/status/1651/problem/B
func TestCF1651B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2
512
3
outputCopy
YES
1 337
NO
YES
31 4 159`
	testutil.AssertEqualCase(t, rawText, 0, CF1651B)
}
