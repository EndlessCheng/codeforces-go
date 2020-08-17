package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1359E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 3
outputCopy
16
inputCopy
3 7
outputCopy
0
inputCopy
1337 42
outputCopy
95147305
inputCopy
1 1
outputCopy
1
inputCopy
500000 1
outputCopy
500000`
	testutil.AssertEqualCase(t, rawText, 0, CF1359E)
}
