package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1034A(t *testing.T) {
	// just copy from website
	rawText := `
3
1 2 4
outputCopy
1
inputCopy
4
6 9 15 30
outputCopy
2
inputCopy
3
1 1 1
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1034A)
}
