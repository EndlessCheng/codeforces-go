package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1257F(t *testing.T) {
	// just copy from website
	rawText := `
2
7 2
outputCopy
1
inputCopy
4
3 17 6 0
outputCopy
5
inputCopy
3
1 2 3
outputCopy
-1
inputCopy
3
43 12 12
outputCopy
1073709057`
	testutil.AssertEqualCase(t, rawText, 0, Sol1257F)
}
