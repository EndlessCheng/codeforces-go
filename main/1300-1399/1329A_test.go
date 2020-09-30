package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1329A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 3
3 2 2
outputCopy
2 4 1
inputCopy
10 1
1
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1329A)
}
