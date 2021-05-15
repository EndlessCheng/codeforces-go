package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/610/C
// https://codeforces.com/problemset/status/610/problem/C
func TestCF610C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
outputCopy
++**
+*+*
++++
+**+`
	testutil.AssertEqualCase(t, rawText, 0, CF610C)
}
