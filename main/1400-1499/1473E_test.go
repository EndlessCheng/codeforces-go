package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1473/problem/E
// https://codeforces.com/problemset/status/1473/problem/E
func TestCF1473E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 4
5 3 4
2 1 1
3 2 2
2 4 2
outputCopy
1 2 2 4 
inputCopy
6 8
3 1 1
3 6 2
5 4 2
4 2 2
6 1 1
5 2 1
3 2 3
1 5 4
outputCopy
2 1 4 3 1 
inputCopy
7 10
7 5 5
2 3 3
4 7 1
5 3 6
2 7 6
6 2 6
3 7 6
4 2 1
3 1 4
1 7 4
outputCopy
3 4 2 7 7 3 `
	testutil.AssertEqualCase(t, rawText, 0, CF1473E)
}
