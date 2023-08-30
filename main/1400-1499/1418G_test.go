package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1418/G
// https://codeforces.com/problemset/status/1418/problem/G
func TestCF1418G(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9
1 2 2 2 1 1 2 2 2
outputCopy
3
inputCopy
10
1 2 3 4 1 2 3 1 2 3
outputCopy
0
inputCopy
12
1 2 3 4 3 4 2 1 3 4 2 1
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1418G)
}
