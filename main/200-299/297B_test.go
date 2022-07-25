package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/297/B
// https://codeforces.com/problemset/status/297/problem/B
func TestCF297B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3 3
2 2 2
1 1 3
outputCopy
YES
inputCopy
4 7 9
5 2 7 3
3 5 2 7 3 8 7
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF297B)
}
