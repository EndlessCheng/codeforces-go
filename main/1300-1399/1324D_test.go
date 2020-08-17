package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1324D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
4 8 2 6 2
4 5 4 1 3
outputCopy
7
inputCopy
4
1 3 2 4
1 3 2 4
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1324D)
}
