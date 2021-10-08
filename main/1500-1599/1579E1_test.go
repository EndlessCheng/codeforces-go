package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1579/problem/E1
// https://codeforces.com/problemset/status/1579/problem/E1
func TestCF1579E1(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
4
3 1 2 4
3
3 2 1
3
3 1 2
2
1 2
2
2 1
outputCopy
1 3 2 4 
1 2 3 
1 3 2 
1 2 
1 2 `
	testutil.AssertEqualCase(t, rawText, 0, CF1579E1)
}
