package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1098A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 1 1 1
1 -1 -1 -1 -1
outputCopy
1
inputCopy
5
1 2 3 1
1 -1 2 -1 -1
outputCopy
2
inputCopy
3
1 2
2 -1 1
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1098A)
}
