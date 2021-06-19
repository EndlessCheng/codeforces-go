package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/429/B
// https://codeforces.com/problemset/status/429/problem/B
func TestCF429B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
100 100 100
100 1 100
100 100 100
outputCopy
800`
	testutil.AssertEqualCase(t, rawText, 0, CF429B)
}
