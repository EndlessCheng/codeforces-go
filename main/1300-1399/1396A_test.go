package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1396A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 3 2 4
outputCopy
1 1 
-1
3 4
4 2
2 4
-3 -6 -6
inputCopy
1
99
outputCopy
1 1
-99
1 1
0
1 1
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1396A)
}
