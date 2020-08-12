package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1037D(t *testing.T) {
	// just copy from website
	rawText := `
4
1 2
1 3
2 4
1 2 3 4
outputCopy
Yes
inputCopy
4
1 2
1 3
2 4
1 2 4 3
outputCopy
No`
	testutil.AssertEqualCase(t, rawText, -1, Sol1037D)
}
