package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/484/D
// https://codeforces.com/problemset/status/484/problem/D
func TestCF484D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2 3 1 2
outputCopy
3
inputCopy
3
3 3 3
outputCopy
0
inputCopy
4
23 7 5 1
outputCopy
22`
	testutil.AssertEqualCase(t, rawText, 0, CF484D)
}
