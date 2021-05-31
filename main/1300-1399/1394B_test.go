package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1394/problem/B
// https://codeforces.com/problemset/status/1394/problem/B
func TestCF1394B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 6 3
4 2 1
1 2 2
2 4 3
4 1 4
4 3 5
3 1 6
outputCopy
2
inputCopy
5 5 1
1 4 1
5 1 2
2 5 3
4 3 4
3 2 5
outputCopy
1
inputCopy
6 13 4
3 5 1
2 5 2
6 3 3
1 4 4
2 6 5
5 3 6
4 1 7
4 3 8
5 2 9
4 2 10
2 1 11
6 1 12
4 6 13
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1394B)
}
