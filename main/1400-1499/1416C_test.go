package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1416C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
0 1 3 2
outputCopy
1 0
inputCopy
9
10 7 9 10 7 5 5 3 5
outputCopy
4 14
inputCopy
3
8 10 3
outputCopy
0 8`
	testutil.AssertEqualCase(t, rawText, 0, CF1416C)
}
