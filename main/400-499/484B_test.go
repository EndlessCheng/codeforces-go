package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/484/B
// https://codeforces.com/problemset/status/484/problem/B
func TestCF484B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3 4 5
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF484B)
}
