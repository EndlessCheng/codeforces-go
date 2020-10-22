package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1364C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2 3
outputCopy
0 1 2 
inputCopy
4
0 0 0 2
outputCopy
1 3 4 0 
inputCopy
3
1 1 3
outputCopy
0 2 1 `
	testutil.AssertEqualCase(t, rawText, 0, CF1364C)
}
