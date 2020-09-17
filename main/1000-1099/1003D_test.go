package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1003D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 4
2 4 8 2 4
8
5
14
10
outputCopy
1
-1
3
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1003D)
}
