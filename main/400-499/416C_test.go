package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/416/C
// https://codeforces.com/problemset/status/416/problem/C
func TestCF416C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
10 50
2 100
5 30
3
4 6 9
outputCopy
2 130
2 1
3 2`
	testutil.AssertEqualCase(t, rawText, 0, CF416C)
}
