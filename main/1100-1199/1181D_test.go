package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1181/D
// https://codeforces.com/problemset/status/1181/problem/D
func TestCF1181D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 4 10
3 1 1 1 2 2
7
8
9
10
11
12
13
14
15
16
outputCopy
4
3
4
2
3
4
1
2
3
4
inputCopy
4 5 4
4 4 5 1
15
9
13
6
outputCopy
5
3
3
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1181D)
}
