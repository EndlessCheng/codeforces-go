package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/986/C
// https://codeforces.com/problemset/status/986/problem/C
func TestCF986C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 3
1 2 3
outputCopy
2
inputCopy
5 5
5 19 10 20 12
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF986C)
}
