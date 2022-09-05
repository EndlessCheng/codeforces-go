package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/484/problem/A
// https://codeforces.com/problemset/status/484/problem/A
func TestCF484A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2
2 4
1 10
outputCopy
1
3
7`
	testutil.AssertEqualCase(t, rawText, 0, CF484A)
}
