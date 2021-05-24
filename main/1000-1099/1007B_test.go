package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1007/B
// https://codeforces.com/problemset/status/1007/problem/B
func TestCF1007B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 1 1
1 6 1
2 2 2
100 100 100
outputCopy
1
4
4
165
inputCopy
1
1 1 1
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1007B)
}
