package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1290/C 2400
// https://codeforces.com/problemset/status/1290/problem/C?friends=on
func Test_cf1290C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 3
0011100
3
1 4 6
3
3 4 7
2
2 3
outputCopy
1
2
3
3
3
3
3
inputCopy
8 6
00110011
3
1 3 8
5
1 2 5 6 7
2
6 8
2
3 5
2
4 7
1
2
outputCopy
1
1
1
1
1
1
4
4
inputCopy
5 3
00011
3
1 2 3
1
4
3
3 4 5
outputCopy
1
1
1
1
1
inputCopy
19 5
1001001001100000110
2
2 3
2
5 6
2
8 9
5
12 13 14 15 16
1
19
outputCopy
0
1
1
1
2
2
2
3
3
3
3
4
4
4
4
4
4
4
5`
	testutil.AssertEqualCase(t, rawText, 0, cf1290C)
}
