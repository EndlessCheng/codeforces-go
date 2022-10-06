package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/754/problem/D
// https://codeforces.com/problemset/status/754/problem/D
func TestCF754D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 2
1 100
40 70
120 130
125 180
outputCopy
31
1 2 
inputCopy
3 2
1 12
15 20
25 30
outputCopy
0
1 2 
inputCopy
5 2
1 10
5 15
14 50
30 70
99 100
outputCopy
21
3 4 `
	testutil.AssertEqualCase(t, rawText, 0, CF754D)
}
