package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1528/problem/D
// https://codeforces.com/problemset/status/1528/problem/D
func TestCF1528D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 4
0 1 1
0 2 3
1 0 1
2 0 1
outputCopy
0 1 2 
1 0 2 
1 2 0 
inputCopy
6 6
0 0 1
1 1 1
2 2 1
3 3 1
4 4 1
5 5 1
outputCopy
0 2 3 3 4 4 
4 0 2 3 3 4 
4 4 0 2 3 3 
3 4 4 0 2 3 
3 3 4 4 0 2 
2 3 3 4 4 0 
inputCopy
4 5
0 1 1
1 3 2
2 2 10
3 0 1
0 0 2
outputCopy
0 1 2 3 
3 0 3 2 
12 13 0 11 
1 2 2 0 
inputCopy
3 3
0 0 1000000000
1 1 1000000000
2 2 1000000000
outputCopy
0 1000000001 1000000002 
1000000002 0 1000000001 
1000000001 1000000002 0 `
	testutil.AssertEqualCase(t, rawText, 0, CF1528D)
}
