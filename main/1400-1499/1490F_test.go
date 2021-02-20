package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1490/F
// https://codeforces.com/problemset/status/1490/problem/F
func TestCF1490F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
6
1 3 2 1 4 2
4
100 100 4 100
8
1 2 3 3 3 2 6 6
outputCopy
2
1
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1490F)
}
