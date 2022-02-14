package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1361/problem/C
// https://codeforces.com/problemset/status/1361/problem/C
func TestCF1361C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
13 11
11 1
3 5
17 1
9 27
outputCopy
3
8 7 9 10 5 6 1 2 3 4 
inputCopy
5
13 11
11 1
3 5
17 1
7 29
outputCopy
2
8 7 10 9 5 6 4 3 2 1 
inputCopy
1
1 1
outputCopy
20
2 1 
inputCopy
8
15 10
13 12
4 6
9 2
0 14
3 8
5 7
1 11
outputCopy
3
12 11 16 15 7 8 2 1 14 13 3 4 5 6 10 9 
inputCopy
12
9 11
7 5
8 11
0 0
3 4
6 5
2 10
9 3
8 2
1 6
7 10
4 1
outputCopy
3
8 7 17 18 14 13 22 21 3 4 12 11 20 19 24 23 10 9 16 15 1 2 6 5 `
	testutil.AssertEqualCase(t, rawText, 0, CF1361C)
}
