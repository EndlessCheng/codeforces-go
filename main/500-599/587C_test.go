package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/587/C
// https://codeforces.com/problemset/status/587/problem/C
func TestCF587C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 4 5
1 3
1 2
1 4
4 5
2 1 4 3
4 5 6
1 5 2
5 5 10
2 3 3
5 3 1
outputCopy
1 3
2 2 3
0
3 1 2 4
1 2`
	testutil.AssertEqualCase(t, rawText, 0, CF587C)
}
