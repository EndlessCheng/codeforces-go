package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1646/problem/B
// https://codeforces.com/problemset/status/1646/problem/B
func TestCF1646B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3
1 2 3
5
2 8 6 3 1
4
3 5 4 2
5
1000000000 1000000000 1000000000 1000000000 1000000000
outputCopy
NO
YES
NO
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1646B)
}
