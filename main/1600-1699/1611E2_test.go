package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1611/problem/E2
// https://codeforces.com/problemset/status/1611/problem/E2
func TestCF1611E2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4

8 2
5 3
4 7
2 5
1 6
3 6
7 2
1 7
6 8

8 4
6 5 7 3
4 7
2 5
1 6
3 6
7 2
1 7
6 8

3 1
2
1 2
2 3

3 2
2 3
3 1
1 2
outputCopy
-1
2
1
2
inputCopy
1
4 1
3
2 4
3 2
2 1
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, -1, CF1611E2)
}
