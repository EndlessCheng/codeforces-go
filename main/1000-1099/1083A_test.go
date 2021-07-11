package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1083/problem/A
// https://codeforces.com/problemset/status/1083/problem/A
func TestCF1083A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 3 3
1 2 2
1 3 2
outputCopy
3
inputCopy
5
6 3 2 5 0
1 2 10
2 3 3
2 4 1
1 5 1
outputCopy
7`
	testutil.AssertEqualCase(t, rawText, 0, CF1083A)
}
