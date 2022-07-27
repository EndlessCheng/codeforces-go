package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/494/problem/A
// https://codeforces.com/problemset/status/494/problem/A
func TestCF494A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
(((#)((#)
outputCopy
1
2
inputCopy
()((#((#(#()
outputCopy
2
2
1
inputCopy
#
outputCopy
-1
inputCopy
(#)
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF494A)
}
