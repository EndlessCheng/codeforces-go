package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1277/D
// https://codeforces.com/problemset/status/1277/problem/D
func TestCF1277D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4
0001
1000
0011
0111
3
010
101
0
2
00000
00001
4
01
001
0001
00001
outputCopy
1
3 
-1
0

2
1 2 
inputCopy
1
5
101
1010
0111
100
110
outputCopy
1
2`
	testutil.AssertEqualCase(t, rawText, -1, CF1277D)
}
