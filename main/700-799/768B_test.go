package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/768/B
// https://codeforces.com/problemset/status/768/problem/B
func TestCF768B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 2 5
outputCopy
4
inputCopy
10 3 10
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, 0, CF768B)
}
