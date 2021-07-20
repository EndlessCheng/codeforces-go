package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/791/D
// https://codeforces.com/problemset/status/791/problem/D
func TestCF791D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 2
1 2
1 3
2 4
2 5
4 6
outputCopy
20
inputCopy
13 3
1 2
3 2
4 2
5 2
3 6
10 6
6 7
6 13
5 8
5 9
9 11
11 12
outputCopy
114
inputCopy
3 5
2 1
3 1
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, -1, CF791D)
}
