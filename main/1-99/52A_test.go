package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/52/A
// https://codeforces.com/problemset/status/52/problem/A
func TestCF52A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9
1 3 2 2 2 1 1 2 3
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, 0, CF52A)
}
