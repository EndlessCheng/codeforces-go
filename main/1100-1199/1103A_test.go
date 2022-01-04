package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1103/A
// https://codeforces.com/problemset/status/1103/problem/A
func TestCF1103A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
010
outputCopy
1 1
1 2
1 4`
	testutil.AssertEqualCase(t, rawText, 0, CF1103A)
}
