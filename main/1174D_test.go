package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1174D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 5
outputCopy
3
6 1 3
inputCopy
2 4
outputCopy
3
1 3 1 
inputCopy
1 1
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1174D)
}
