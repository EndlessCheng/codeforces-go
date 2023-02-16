package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1579/problem/B
// https://codeforces.com/problemset/status/1579/problem/B
func TestCF1579B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2
2 1
3
1 2 1
4
2 4 1 3
5
2 5 1 4 3
outputCopy
1
1 2 1
1
1 3 2
3
2 4 1
2 3 1
1 3 2
4
2 4 2
1 5 3
1 2 1
1 3 1`
	testutil.AssertEqualCase(t, rawText, 0, CF1579B)
}
