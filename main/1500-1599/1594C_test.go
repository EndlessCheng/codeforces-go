package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1594/problem/C
// https://codeforces.com/problemset/status/1594/problem/C
func TestCF1594C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4 a
aaaa
4 a
baaa
4 b
bzyx
outputCopy
0
1
2
2 
2 3`
	testutil.AssertEqualCase(t, rawText, 0, CF1594C)
}
