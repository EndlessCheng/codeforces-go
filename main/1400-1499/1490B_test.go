package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1490/B
// https://codeforces.com/problemset/status/1490/problem/B
func TestCF1490B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
6
0 2 5 5 4 8
6
2 0 2 1 0 0
9
7 1 3 4 2 10 3 9 6
6
0 1 2 3 4 5
outputCopy
3
1
3
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1490B)
}
