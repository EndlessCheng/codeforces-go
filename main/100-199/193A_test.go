package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/193/A
// https://codeforces.com/problemset/status/193/problem/A
func TestCF193A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 4
####
#..#
#..#
#..#
####
outputCopy
2
inputCopy
5 5
#####
#...#
#####
#...#
#####
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF193A)
}
