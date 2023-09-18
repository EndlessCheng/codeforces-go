package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1872/problem/E
// https://codeforces.com/problemset/status/1872/problem/E
func TestCF1872E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5
1 2 3 4 5
01000
7
2 0
2 1
1 2 4
2 0
2 1
1 1 3
2 1
6
12 12 14 14 5 5
001001
3
2 1
1 2 4
2 1
4
7 7 7 777
1111
3
2 0
1 2 3
2 0
2
1000000000 996179179
11
1
2 1
5
1 42 20 47 7
00011
5
1 3 4
1 1 1
1 3 4
1 2 4
2 0
outputCopy
3 2 6 7 7 
11 7 
0 0 
16430827 
47 `
	testutil.AssertEqualCase(t, rawText, 0, CF1872E)
}
