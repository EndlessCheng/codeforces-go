package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/289/B
// https://codeforces.com/problemset/status/289/problem/B
func TestCF289B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 2 2
2 4
6 8
outputCopy
4
inputCopy
1 2 7
6 7
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF289B)
}
