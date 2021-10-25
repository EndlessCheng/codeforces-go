package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1512/F
// https://codeforces.com/problemset/status/1512/problem/F
func TestCF1512F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4 15
1 3 10 11
1 2 7
4 100
1 5 10 50
3 14 12
2 1000000000
1 1
1
outputCopy
6
13
1000000000`
	testutil.AssertEqualCase(t, rawText, 0, CF1512F)
}
