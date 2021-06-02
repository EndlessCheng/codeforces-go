package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1282/problem/E
// https://codeforces.com/problemset/status/1282/problem/E
func TestCF1282E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
6
3 6 5
5 2 4
5 4 6
6 3 1
6
2 5 6
2 5 1
4 1 2
1 3 5
3
1 2 3
outputCopy
1 6 4 2 5 3 
4 2 3 1 
1 4 2 6 5 3 
3 4 2 1 
1 3 2 
1 `
	testutil.AssertEqualCase(t, rawText, 0, CF1282E)
}
