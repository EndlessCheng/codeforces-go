package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/435/B
// https://codeforces.com/problemset/status/435/problem/B
func TestCF435B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1990 1
outputCopy
9190
inputCopy
300 0
outputCopy
300
inputCopy
1034 2
outputCopy
3104
inputCopy
9090000078001234 6
outputCopy
9907000008001234`
	testutil.AssertEqualCase(t, rawText, 0, CF435B)
}
