package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1474/problem/C
// https://codeforces.com/problemset/status/1474/problem/C
func TestCF1474C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2
3 5 1 2
3
1 1 8 8 64 64
2
1 1 2 4
5
1 2 3 4 5 6 7 14 3 11
outputCopy
YES
6
1 5
2 3
NO
NO
YES
21
14 7
3 11
5 6
2 4
3 1`
	testutil.AssertEqualCase(t, rawText, -1, CF1474C)
}
