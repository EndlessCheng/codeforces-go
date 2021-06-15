package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1156/E
// https://codeforces.com/problemset/status/1156/problem/E
func TestCF1156E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3 4 1 5 2
outputCopy
2
inputCopy
3
1 3 2
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1156E)
}
