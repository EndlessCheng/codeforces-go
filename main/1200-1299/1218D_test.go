package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1218/problem/D
// https://codeforces.com/problemset/status/1218/problem/D
func TestCF1218D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 6
4 1 5
5 2 1
6 3 2
1 2 6
1 3 3
2 3 4
outputCopy
1 1`
	testutil.AssertEqualCase(t, rawText, 0, CF1218D)
}
