package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1467/A
// https://codeforces.com/problemset/status/1467/problem/A
func TestCF1467A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1
2
outputCopy
9
98`
	testutil.AssertEqualCase(t, rawText, 0, CF1467A)
}
