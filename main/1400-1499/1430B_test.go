package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1430B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
4 1
5 5 5 5
3 2
0 0 0
outputCopy
10
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1430B)
}
