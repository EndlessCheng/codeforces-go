package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1328C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
5
22222
5
21211
1
2
9
220222021
outputCopy
11111
11111
11000
10211
1
1
110111011
110111010`
	testutil.AssertEqualCase(t, rawText, 0, CF1328C)
}
