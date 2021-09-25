package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1473/C
// https://codeforces.com/problemset/status/1473/problem/C
func TestCF1473C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 1
2 2
3 2
4 3
outputCopy
1 
1 2 
2 1 
1 3 2 `
	testutil.AssertEqualCase(t, rawText, 0, CF1473C)
}
