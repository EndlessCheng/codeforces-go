package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF803C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 3
outputCopy
1 2 3
inputCopy
8 2
outputCopy
2 6
inputCopy
5 3
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF803C)
}
