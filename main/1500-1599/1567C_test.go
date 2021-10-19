package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1567/C
// https://codeforces.com/problemset/status/1567/problem/C
func TestCF1567C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
100
12
8
2021
10000
outputCopy
9
4
7
44
99`
	testutil.AssertEqualCase(t, rawText, 0, CF1567C)
}
