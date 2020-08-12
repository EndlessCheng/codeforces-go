package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1009C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 3
-1 3
0 0
-1 -4
outputCopy
-2.500000000000000
inputCopy
3 2
0 2
5 0
outputCopy
7.000000000000000`
	testutil.AssertEqualCase(t, rawText, 0, CF1009C)
}
