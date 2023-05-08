package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1753/C
// https://codeforces.com/problemset/status/1753/problem/C
func TestCF1753C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3
0 1 0
5
0 0 1 1 1
6
1 1 1 0 0 1
outputCopy
3
0
249561107`
	testutil.AssertEqualCase(t, rawText, 0, CF1753C)
}
