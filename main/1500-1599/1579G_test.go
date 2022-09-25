package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1579/G
// https://codeforces.com/problemset/status/1579/problem/G
func TestCF1579G(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
2
1 3
3
1 2 3
4
6 2 3 9
4
6 8 4 5
7
1 2 4 6 7 7 3
8
8 6 5 1 2 2 3 6
outputCopy
3
3
9
9
7
8`
	testutil.AssertEqualCase(t, rawText, 0, CF1579G)
}
