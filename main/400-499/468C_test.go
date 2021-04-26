package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/468/C
// https://codeforces.com/problemset/status/468/problem/C
func TestCF468C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
46
outputCopy
1 10
inputCopy
126444381000032
outputCopy
2333333 2333333333333`
	testutil.AssertEqualCase(t, rawText, 0, CF468C)
}
