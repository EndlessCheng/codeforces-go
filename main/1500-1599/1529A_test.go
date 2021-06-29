package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1529/A
// https://codeforces.com/problemset/status/1529/problem/A
func TestCF1529A(t *testing.T) {
	// just copy from website
	rawText := `
3
6
1 1 1 2 2 3
6
9 9 9 9 9 9
6
6 4 1 1 4 1
outputCopy
3
0
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1529A)
}
