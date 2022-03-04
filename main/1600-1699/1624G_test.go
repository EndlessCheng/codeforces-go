package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1624/problem/G
// https://codeforces.com/problemset/status/1624/problem/G
func TestCF1624G(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3

3 3
1 2 1
2 3 2
1 3 2

5 7
4 2 7
2 5 8
3 4 2
3 2 1
2 4 2
4 1 2
1 2 2

3 4
1 2 1
2 3 2
1 3 3
3 1 4
outputCopy
2
10
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1624G)
}
