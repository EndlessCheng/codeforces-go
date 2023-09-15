package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1322/B
// https://codeforces.com/problemset/status/1322/problem/B
func TestCF1322B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 2
outputCopy
3
inputCopy
3
1 2 3
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1322B)
}
