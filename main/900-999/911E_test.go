package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/911/problem/E
// https://codeforces.com/problemset/status/911/problem/E
func TestCF911E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 3
3 2 1
outputCopy
3 2 1 5 4 
inputCopy
5 3
2 3 1
outputCopy
-1
inputCopy
5 1
3
outputCopy
3 2 1 5 4 
inputCopy
5 2
3 4
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF911E)
}
