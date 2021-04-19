package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1268/B
// https://codeforces.com/problemset/status/1268/problem/B
func TestCF1268B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3 2 2 2 1
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1268B)
}
