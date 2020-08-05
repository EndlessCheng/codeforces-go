package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol494A(t *testing.T) {
	// just copy from website
	rawText := `
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
	testutil.AssertEqual(t, rawText, Sol494A)
}
