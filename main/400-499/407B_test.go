package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/407/B
// https://codeforces.com/problemset/status/407/problem/B
func TestCF407B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 2
outputCopy
4
inputCopy
4
1 1 2 3
outputCopy
20
inputCopy
5
1 1 1 1 1
outputCopy
62`
	testutil.AssertEqualCase(t, rawText, 0, CF407B)
}
