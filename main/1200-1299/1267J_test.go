package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1267/problem/J
// https://codeforces.com/problemset/status/1267/problem/J
func TestCF1267J(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
11
1 5 1 5 1 5 1 1 1 1 5
6
1 2 2 2 2 1
5
4 3 3 1 2
outputCopy
3
3
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1267J)
}
