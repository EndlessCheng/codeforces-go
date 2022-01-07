package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/920/problem/D
// https://codeforces.com/problemset/status/920/problem/D
func TestCF920D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 3 5
2 3
outputCopy
YES
1 2 1
inputCopy
2 3 4
2 3
outputCopy
NO
inputCopy
5 2 0
1 3 5 7 9
outputCopy
YES
2 2 1
3 3 1
4 4 1
5 5 1
inputCopy
6 4 8
5 5 5 5 5 5
outputCopy
YES
2 2 1
2 3 1
2 4 1
2 5 1
2 6 1
2 1 6
inputCopy
6 11 3
6 6 6 6 6 6
outputCopy
YES
1 2 1
1 3 1
1 4 1
1 5 1
1 6 1
3 1 6
inputCopy
2 4 4
2 3
outputCopy
YES
1 2 1
1 1 2`
	testutil.AssertEqualCase(t, rawText, 0, CF920D)
}
