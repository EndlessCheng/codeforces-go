package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1628/A
// https://codeforces.com/problemset/status/1628/problem/A
func TestCF1628A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
5
1 0 2 0 3
8
2 2 3 4 0 1 2 0
1
1
5
0 1 2 3 4
4
0 1 1 0
10
0 0 2 1 1 1 0 0 1 1
outputCopy
1
4 
2
5 1 
1
0 
1
5 
2
2 2 
4
3 2 2 0 `
	testutil.AssertEqualCase(t, rawText, 0, CF1628A)
}
