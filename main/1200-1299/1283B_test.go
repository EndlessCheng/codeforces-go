package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1283/problem/B
// https://codeforces.com/problemset/status/1283/problem/B
func TestCF1283B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5 2
19 4
12 7
6 2
100000 50010
outputCopy
5
18
10
6
75015`
	testutil.AssertEqualCase(t, rawText, 0, CF1283B)
}
