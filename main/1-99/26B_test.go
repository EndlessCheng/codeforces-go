package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/26/B
// https://codeforces.com/problemset/status/26/problem/B
func TestCF26B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
(()))(
outputCopy
4
inputCopy
((()())
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF26B)
}
