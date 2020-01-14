package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF676D(t *testing.T) {
	// just copy from website
	rawText := `
2 2
+*
*U
1 1
2 2
outputCopy
-1
inputCopy
2 3
<><
><>
1 1
2 1
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF676D)
}
