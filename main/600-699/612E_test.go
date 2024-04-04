package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/612/E
// https://codeforces.com/problemset/status/612/problem/E
func TestCF612E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2 1 4 3
outputCopy
3 4 2 1
inputCopy
4
2 1 3 4
outputCopy
-1
inputCopy
5
2 3 4 5 1
outputCopy
4 5 1 2 3`
	testutil.AssertEqualCase(t, rawText, -1, CF612E)
}
