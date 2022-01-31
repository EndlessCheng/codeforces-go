package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1632/C
// https://codeforces.com/problemset/status/1632/problem/C
func TestCF1632C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 3
5 8
2 5
3 19
56678 164422
outputCopy
1
3
2
1
23329`
	testutil.AssertEqualCase(t, rawText, 0, CF1632C)
}
