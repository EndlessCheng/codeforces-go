package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1771/D
// https://codeforces.com/problemset/status/1771/problem/D
func TestCF1771D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
5
abaca
1 2
1 3
3 4
4 5
9
caabadedb
1 2
2 3
2 4
1 5
5 6
5 7
5 8
8 9
outputCopy
3
5`
	testutil.AssertEqualCase(t, rawText, 0, CF1771D)
}
