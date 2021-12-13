package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1455/B
// https://codeforces.com/problemset/status/1455/problem/B
func TestCF1455B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1
2
3
4
5
outputCopy
1
3
2
3
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1455B)
}
