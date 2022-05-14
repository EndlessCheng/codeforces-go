package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1338/A
// https://codeforces.com/problemset/status/1338/problem/A
func TestCF1338A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4
1 7 6 5
5
1 2 3 4 5
2
0 -4
outputCopy
2
0
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1338A)
}
