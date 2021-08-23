package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1228/E
// https://codeforces.com/problemset/status/1228/problem/E
func TestCF1228E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 2
outputCopy
7
inputCopy
123 456789
outputCopy
689974806`
	testutil.AssertEqualCase(t, rawText, 0, CF1228E)
}
