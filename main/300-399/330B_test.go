package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/330/B
// https://codeforces.com/problemset/status/330/problem/B
func TestCF330B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 1
1 3
outputCopy
3
1 2
4 2
2 3`
	testutil.AssertEqualCase(t, rawText, 0, CF330B)
}
