package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1593/E
// https://codeforces.com/problemset/status/1593/problem/E
func TestCF1593E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6

14 1
1 2
2 3
2 4
4 5
4 6
2 7
7 8
8 9
8 10
3 11
3 12
1 13
13 14

2 200000
1 2

3 2
1 2
2 3

5 1
5 1
3 2
2 1
5 4

6 2
5 1
2 5
5 6
4 2
3 4

7 1
4 3
5 1
1 3
6 1
1 7
2 1
outputCopy
7
0
0
3
1
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1593E)
}
