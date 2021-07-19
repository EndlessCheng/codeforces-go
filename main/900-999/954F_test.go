package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/954/F
// https://codeforces.com/problemset/status/954/problem/F
func TestCF954F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 5
1 3 4
2 2 3
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF954F)
}
