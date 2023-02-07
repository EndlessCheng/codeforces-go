package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1468/D
// https://codeforces.com/problemset/status/1468/problem/D
func TestCF1468D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
7 2 3 6
1 4
7 2 3 6
5 1
7 2 3 6
4 4
outputCopy
2
1
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1468D)
}
