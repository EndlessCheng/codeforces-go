package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1176D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3 5 2 3 2 4
outputCopy
3 4 2 
inputCopy
1
2750131 199999
outputCopy
199999 
inputCopy
1
3 6
outputCopy
6 
inputCopy
1
815 1630
outputCopy
1630`
	testutil.AssertEqualCase(t, rawText, -1, CF1176D)
}
