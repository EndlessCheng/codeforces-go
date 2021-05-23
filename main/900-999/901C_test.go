package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/901/C
// https://codeforces.com/problemset/status/901/problem/C
func TestCF901C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 6
1 2
2 3
3 1
4 5
5 6
6 4
3
1 3
4 6
1 6
outputCopy
5
5
14
inputCopy
8 9
1 2
2 3
3 1
4 5
5 6
6 7
7 8
8 4
7 2
3
1 8
1 4
3 8
outputCopy
27
8
19`
	testutil.AssertEqualCase(t, rawText, 0, CF901C)
}
