package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1419/D2
// https://codeforces.com/problemset/status/1419/problem/D2
func TestCF1419D2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
1 3 2 2 4 5 4
outputCopy
3
3 1 4 2 4 2 5 `
	testutil.AssertEqualCase(t, rawText, 0, CF1419D2)
}
