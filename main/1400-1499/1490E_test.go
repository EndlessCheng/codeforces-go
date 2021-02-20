package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1490/E
// https://codeforces.com/problemset/status/1490/problem/E
func TestCF1490E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
4
1 2 4 3
5
1 1 1 1 1
outputCopy
3
2 3 4 
5
1 2 3 4 5 `
	testutil.AssertEqualCase(t, rawText, 0, CF1490E)
}
