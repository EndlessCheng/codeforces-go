package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/696/C
// https://codeforces.com/problemset/status/696/problem/C
func TestCF696C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
2
outputCopy
1/2
inputCopy
3
1 1 1
outputCopy
0/1`
	testutil.AssertEqualCase(t, rawText, 0, CF696C)
}
