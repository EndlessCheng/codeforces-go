package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1179/B
// https://codeforces.com/problemset/status/1179/problem/B
func TestCF1179B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 3
outputCopy
1 1
1 3
1 2
2 2
2 3
2 1
inputCopy
1 1
outputCopy
1 1
inputCopy
3 4
outputCopy
`
	testutil.AssertEqualCase(t, rawText, 0, CF1179B)
}
