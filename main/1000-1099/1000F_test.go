package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1000F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
1 1 2 3 2 4
2
2 6
1 2
outputCopy
4
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1000F)
}
