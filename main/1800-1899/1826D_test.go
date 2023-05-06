package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1826/problem/D
// https://codeforces.com/problemset/status/1826/problem/D
func TestCF1826D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
5
5 1 4 2 3
4
1 1 1 1
6
9 8 7 6 5 4
7
100000000 1 100000000 1 100000000 1 100000000
outputCopy
8
1
22
299999996`
	testutil.AssertEqualCase(t, rawText, 0, CF1826D)
}
