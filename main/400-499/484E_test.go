package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/484/E
// https://codeforces.com/problemset/status/484/problem/E
func TestCF484E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2 2 3 3
3
2 5 3
2 5 2
1 5 5
outputCopy
2
3
1`
	testutil.AssertEqualCase(t, rawText, 0, CF484E)
}
