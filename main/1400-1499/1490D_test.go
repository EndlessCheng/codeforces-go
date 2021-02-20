package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1490/D
// https://codeforces.com/problemset/status/1490/problem/D
func TestCF1490D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
5
3 5 2 1 4
1
1
4
4 3 1 2
outputCopy
1 0 2 3 1 
0 
0 1 3 2 `
	testutil.AssertEqualCase(t, rawText, 0, CF1490D)
}
