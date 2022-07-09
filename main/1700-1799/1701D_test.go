package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1701/problem/D
// https://codeforces.com/problemset/status/1701/problem/D
func TestCF1701D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4
0 2 0 1
2
1 1
5
0 0 1 4 1
3
0 1 3
outputCopy
2 1 4 3 
1 2 
3 4 2 1 5 
3 2 1
inputCopy
1
21
0 0 0 0 1 0 0 8 0 0 0 0 2 1 0 0 1 2 2 5 10
outputCopy
6 11 12 13 3 14 15 1 16 17 18 19 5 9 20 21 10 7 8 4 2`
	testutil.AssertEqualCase(t, rawText, 0, CF1701D)
}
