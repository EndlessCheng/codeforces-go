package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1225/D
// https://codeforces.com/problemset/status/1225/problem/D
func TestCF1225D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 3
1 3 9 8 24 1
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, 1, CF1225D)
}
