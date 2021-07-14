package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1547/F
// https://codeforces.com/problemset/status/1547/problem/F
func TestCF1547F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
4
16 24 10 5
4
42 42 42 42
3
4 6 4
5
1 2 3 4 5
6
9 9 27 9 9 63
outputCopy
3
0
2
1
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1547F)
}
