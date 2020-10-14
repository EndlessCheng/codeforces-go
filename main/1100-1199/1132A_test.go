package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1132A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1
4
3
outputCopy
1
inputCopy
0
0
0
0
outputCopy
1
inputCopy
1
2
3
4
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1132A)
}
