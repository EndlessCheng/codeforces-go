package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1618/problem/E
// https://codeforces.com/problemset/status/1618/problem/E
func TestCF1618E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3
12 16 14
1
1
3
1 2 3
6
81 75 75 93 93 87
outputCopy
YES
3 1 3 
YES
1 
NO
YES
5 5 4 1 4 5 `
	testutil.AssertEqualCase(t, rawText, 0, CF1618E)
}
