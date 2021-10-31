package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1594/B
// https://codeforces.com/problemset/status/1594/problem/B
func TestCF1594B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3 4
2 12
105 564
outputCopy
9
12
3595374`
	testutil.AssertEqualCase(t, rawText, 0, CF1594B)
}
