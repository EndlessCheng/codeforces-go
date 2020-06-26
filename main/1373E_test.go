package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1373E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9
1 0
1 1
42 7
13 7
99 1
99 0
99 2
149 1
150 1
outputCopy
1
0
4
-1
599998
99999999999
7997`
	testutil.AssertEqualCase(t, rawText, 0, CF1373E)
}
