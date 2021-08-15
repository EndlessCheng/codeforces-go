package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/584/E
// https://codeforces.com/problemset/status/584/problem/E
func TestCF584E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4 2 1 3
3 2 4 1
outputCopy
3
2
4 3
3 1`
	testutil.AssertEqualCase(t, rawText, 0, CF584E)
}
