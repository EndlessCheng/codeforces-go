package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/466/D
// https://codeforces.com/problemset/status/466/problem/D
func TestCF466D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
1 1 1
outputCopy
4
inputCopy
5 1
1 1 1 1 1
outputCopy
1
inputCopy
4 3
3 2 1 1
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF466D)
}
