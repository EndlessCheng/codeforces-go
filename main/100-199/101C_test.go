package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF101C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
0 0
1 1
0 1
outputCopy
YES
inputCopy
0 0
1 1
1 1
outputCopy
YES
inputCopy
0 0
1 1
2 2
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF101C)
}
