package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1148/C
// https://codeforces.com/problemset/status/1148/problem/C
func TestCF1148C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
2 1
outputCopy
1
1 2
inputCopy
4
3 4 1 2
outputCopy
4
1 4
1 4
1 3
2 4
inputCopy
6
2 5 3 1 4 6
outputCopy
3
1 5
2 5
1 4`
	testutil.AssertEqualCase(t, rawText, 0, CF1148C)
}
