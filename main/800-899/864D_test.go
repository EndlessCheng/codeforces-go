package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/864/D
// https://codeforces.com/problemset/status/864/problem/D
func TestCF864D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3 2 2 3
outputCopy
2
1 2 4 3 
inputCopy
6
4 5 6 3 2 1
outputCopy
0
4 5 6 3 2 1 
inputCopy
10
6 8 4 6 7 1 6 3 4 5
outputCopy
3
2 8 4 6 7 1 9 3 10 5 `
	testutil.AssertEqualCase(t, rawText, 0, CF864D)
}
