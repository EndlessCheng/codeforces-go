package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/245/H
// https://codeforces.com/problemset/status/245/problem/H
func TestCF245H(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
caaaba
5
1 1
1 4
2 3
4 6
4 5
outputCopy
1
7
3
4
2`
	testutil.AssertEqualCase(t, rawText, 0, CF245H)
}
