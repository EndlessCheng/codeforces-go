package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1283/F
// https://codeforces.com/problemset/status/1283/problem/F
func TestCF1283F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
3 6 3 1 5
outputCopy
3
6 3
6 5
1 3
1 4
5 2`
	testutil.AssertEqualCase(t, rawText, 0, CF1283F)
}
