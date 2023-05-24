package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1833/E
// https://codeforces.com/problemset/status/1833/problem/E
func TestCF1833E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10
6
2 1 4 3 6 5
6
2 3 1 5 6 4
9
2 3 2 5 6 5 8 9 8
2
2 1
4
4 3 2 1
5
2 3 4 5 1
6
5 3 4 1 1 2
5
3 5 4 1 2
6
6 3 2 5 4 3
6
5 1 4 3 4 2
outputCopy
1 3
2 2
1 3
1 1
1 2
1 1
1 1
2 2
1 2
1 1`
	testutil.AssertEqualCase(t, rawText, 0, CF1833E)
}
