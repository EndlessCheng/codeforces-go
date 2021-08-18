package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/718/C
// https://codeforces.com/problemset/status/718/problem/C
func TestCF718C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 4
1 1 2 1 1
2 1 5
1 2 4 2
2 2 4
2 1 5
outputCopy
5
7
9`
	testutil.AssertEqualCase(t, rawText, 0, CF718C)
}
