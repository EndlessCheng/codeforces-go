package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1539/C
// https://codeforces.com/problemset/status/1539/problem/C
func TestCF1539C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8 2 3
1 1 5 8 12 13 20 22
outputCopy
2
inputCopy
13 0 37
20 20 80 70 70 70 420 5 1 5 1 60 90
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 1, CF1539C)
}
