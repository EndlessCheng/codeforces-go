package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/78/C
// https://codeforces.com/problemset/status/78/problem/C
func TestCF78C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 15 4
outputCopy
Timur
inputCopy
4 9 5
outputCopy
Marsel`
	testutil.AssertEqualCase(t, rawText, 0, CF78C)
}
