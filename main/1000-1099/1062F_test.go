package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1062/problem/F
// https://codeforces.com/problemset/status/1062/problem/F
func TestCF1062F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 7
1 2
2 3
3 4
4 7
2 5
5 4
6 4
outputCopy
4
inputCopy
6 7
1 2
2 3
3 4
1 5
5 3
2 6
6 4
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1062F)
}
