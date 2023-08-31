package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1763/C
// https://codeforces.com/problemset/status/1763/problem/C
func TestCF1763C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3
1 1 1
2
9 1
3
4 9 5
outputCopy
3
16
18`
	testutil.AssertEqualCase(t, rawText, 0, CF1763C)
}
