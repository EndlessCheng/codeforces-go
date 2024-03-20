package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1187/status/E
func TestCF1187E(t *testing.T) {
	// just copy from website
	rawText := `
input
9
1 2
2 3
2 5
2 6
1 4
4 9
9 7
9 8
output
36
input
5
1 2
1 3
2 4
2 5
output
14`
	testutil.AssertEqualCase(t, rawText, 0, CF1187E)
}
