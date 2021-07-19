package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1185/G1
// https://codeforces.com/problemset/status/1185/problem/G1
func TestCF1185G1(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
1 1
1 2
1 3
outputCopy
6
inputCopy
3 3
1 1
1 1
1 3
outputCopy
2
inputCopy
4 10
5 3
2 1
3 2
5 1
outputCopy
10`
	testutil.AssertEqualCase(t, rawText, 0, CF1185G1)
}
