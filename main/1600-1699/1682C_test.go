package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1682/C
// https://codeforces.com/problemset/status/1682/problem/C
func TestCF1682C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3
6 6 6
6
2 5 4 5 2 4
4
1 3 2 2
outputCopy
1
3
2
inputCopy
1
3
1 2 2
outputCopy
2
inputCopy
1
5
1 1 2 3 3
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1682C)
}
