package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/883/K
// https://codeforces.com/problemset/status/883/problem/K
func TestCF883K(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4 5
4 5
4 10
outputCopy
16
9 9 10 
inputCopy
4
1 100
100 1
1 100
100 1
outputCopy
202
101 101 101 101 
inputCopy
3
1 1
100 100
1 1
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF883K)
}
