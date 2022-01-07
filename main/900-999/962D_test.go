package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/962/problem/D
// https://codeforces.com/problemset/status/962/problem/D
func TestCF962D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
3 4 1 2 2 1 1
outputCopy
4
3 8 2 1 
inputCopy
5
1 1 3 1 1
outputCopy
2
3 4 
inputCopy
5
10 40 20 50 30
outputCopy
5
10 40 20 50 30 `
	testutil.AssertEqualCase(t, rawText, 0, CF962D)
}
