package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1266C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 2
outputCopy
4 12
2 9
inputCopy
1 1
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1266C)
}
