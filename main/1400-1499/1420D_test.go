package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1420/D
// https://codeforces.com/problemset/status/1420/problem/D
func TestCF1420D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 3
1 7
3 8
4 5
6 7
1 3
5 10
8 9
outputCopy
9
inputCopy
3 1
1 1
2 2
3 3
outputCopy
3
inputCopy
3 2
1 1
2 2
3 3
outputCopy
0
inputCopy
3 3
1 3
2 3
3 3
outputCopy
1
inputCopy
5 2
1 3
2 4
3 5
4 6
5 7
outputCopy
7
inputCopy
10 3
1 10
2 10
3 10
4 10
5 10
1 2
1 3
1 4
1 5
1 6
outputCopy
80`
	testutil.AssertEqualCase(t, rawText, -1, CF1420D)
}
