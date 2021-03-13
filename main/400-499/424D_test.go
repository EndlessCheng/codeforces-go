package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/424/D
// https://codeforces.com/problemset/status/424/problem/D
func TestCF424D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 7 48
3 6 2
5 4 8 3 3 7 9
4 1 6 8 7 1 1
1 6 4 6 4 8 6
7 2 6 1 6 9 4
1 9 8 6 3 9 2
4 5 6 8 4 3 7
outputCopy
4 3 6 7`
	testutil.AssertEqualCase(t, rawText, 0, CF424D)
}
