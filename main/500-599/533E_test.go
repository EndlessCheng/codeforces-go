package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF533E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
reading
trading
outputCopy
1
inputCopy
5
sweet
sheep
outputCopy
0
inputCopy
3
toy
try
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF533E)
}
