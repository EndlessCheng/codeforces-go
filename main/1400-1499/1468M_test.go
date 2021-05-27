package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1468/problem/M
// https://codeforces.com/problemset/status/1468/problem/M
func TestCF1468M(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4
2 1 10
3 1 3 5
5 5 4 3 2 1
3 10 20 30
3
4 1 2 3 4
4 2 3 4 5
4 3 4 5 6
2
3 1 3 5
3 4 3 2
outputCopy
2 3 
1 2 
-1
inputCopy
1
2
3 1 3 5
3 4 3 2
outputCopy
-1
inputCopy
7
2
2 1 2
2 2 1
3
4 4 5 6 7
3 1 2 3
3 2 3 4
3
3 1 2 3
3 2 3 4
3 3 4 5
2
3 10 20 30
2 40 50
4
2 500 100
2 500 100
2 500 100
2 499 100
2
3 1 3 2
2 1 4
2
2 1 5
2 3 2
outputCopy
1 2 
2 3 
1 2 
-1
1 2 
-1
-1`
	testutil.AssertEqualCase(t, rawText, -1, CF1468M)
}
