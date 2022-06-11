package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1082/B
// https://codeforces.com/problemset/status/1082/problem/B
func TestCF1082B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10
GGGSGGGSGG
outputCopy
7
inputCopy
4
GGGG
outputCopy
4
inputCopy
3
SSS
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1082B)
}
