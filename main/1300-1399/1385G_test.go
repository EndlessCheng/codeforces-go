package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1385/problem/G
// https://codeforces.com/problemset/status/1385/problem/G
func TestCF1385G(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
4
1 2 3 4
2 3 1 4
5
5 3 5 1 4
1 2 3 2 4
3
1 2 1
3 3 2
4
1 2 2 1
3 4 3 4
4
4 3 1 4
3 2 2 1
3
1 1 2
3 2 2
outputCopy
0

2
2 3 
1
1 
2
3 4 
2
3 4 
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1385G)
}
