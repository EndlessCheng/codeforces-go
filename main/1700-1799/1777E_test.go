package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1777/E
// https://codeforces.com/problemset/status/1777/problem/E
func TestCF1777E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2 1
1 2 3
5 4
1 2 10
2 3 10
3 1 10
4 5 10
4 5
1 2 10000
2 3 20000
1 3 30000
4 2 500
4 3 20
4 5
1 2 10000
2 3 20000
1 3 30000
4 2 5
4 3 20
outputCopy
0
-1
20
5
inputCopy
1
6 7
6 4 58
6 1 48
4 1 3
5 6 80
2 3 57
3 1 86
1 2 43
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, -1, CF1777E)
}
