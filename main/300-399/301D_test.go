package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/301/D
// https://codeforces.com/problemset/status/301/problem/D
func TestCF301D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1
1
1 1
outputCopy
1
inputCopy
10 9
1 2 3 4 5 6 7 8 9 10
1 10
2 9
3 8
4 7
5 6
2 2
9 10
5 10
4 10
outputCopy
27
14
8
4
2
1
2
7
9`
	testutil.AssertEqualCase(t, rawText, 0, CF301D)
}
