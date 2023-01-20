package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1598/D
// https://codeforces.com/problemset/status/1598/problem/D
func TestCF1598D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
4
2 4
3 4
2 1
1 3
5
1 5
2 4
3 3
4 2
5 1
outputCopy
3
10`
	testutil.AssertEqualCase(t, rawText, 0, CF1598D)
}
