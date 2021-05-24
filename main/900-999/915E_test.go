package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/915/problem/E
// https://codeforces.com/problemset/status/915/problem/E
func TestCF915E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
6
1 2 1
3 4 1
2 3 2
1 3 2
2 4 1
1 4 2
outputCopy
2
0
2
3
1
4`
	testutil.AssertEqualCase(t, rawText, 0, CF915E)
}
