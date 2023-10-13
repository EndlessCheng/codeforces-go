package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1792/C
// https://codeforces.com/problemset/status/1792/problem/C
func TestCF1792C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
5
1 5 4 2 3
3
1 2 3
4
2 1 4 3
6
5 2 4 1 6 3
outputCopy
2
0
1
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1792C)
}
