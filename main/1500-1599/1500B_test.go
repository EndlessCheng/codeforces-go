package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1500/B
// https://codeforces.com/problemset/status/1500/problem/B
func TestCF1500B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 2 4
4 2 3 1
2 1
outputCopy
5
inputCopy
3 8 41
1 3 2
1 6 4 3 5 7 2 8
outputCopy
47
inputCopy
1 2 31
1
1 2
outputCopy
62
inputCopy
1 2 1
1
2 1
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, -1, CF1500B)
}
