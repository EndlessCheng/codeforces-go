package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1512/E
// https://codeforces.com/problemset/status/1512/problem/E
func TestCF1512E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5 2 3 5
5 3 4 1
3 1 2 4
2 2 2 2
2 1 1 3
outputCopy
1 2 3 4 5 
-1
1 3 2 
1 2 
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1512E)
}
