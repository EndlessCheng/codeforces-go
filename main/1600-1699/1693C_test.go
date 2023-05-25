package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1693/C
// https://codeforces.com/problemset/status/1693/problem/C
func TestCF1693C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 1
1 2
outputCopy
1
inputCopy
4 4
1 2
1 4
2 4
1 4
outputCopy
2
inputCopy
5 7
1 2
2 3
3 5
1 4
4 3
4 5
3 1
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1693C)
}
