package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/460/problem/B
// https://codeforces.com/problemset/status/460/problem/B
func TestCF460B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2 8
outputCopy
3
10 2008 13726 
inputCopy
1 2 -18
outputCopy
0
inputCopy
2 2 -1
outputCopy
4
1 31 337 967 `
	testutil.AssertEqualCase(t, rawText, 0, CF460B)
}
