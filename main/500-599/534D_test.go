package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/534/D
// https://codeforces.com/problemset/status/534/problem/D
func TestCF534D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 1 3 0 1
outputCopy
Possible
4 5 1 3 2 
inputCopy
9
0 2 3 4 1 1 0 2 2
outputCopy
Possible
7 5 2 1 6 8 3 4 9
inputCopy
4
0 2 1 1
outputCopy
Impossible`
	testutil.AssertEqualCase(t, rawText, 0, CF534D)
}
