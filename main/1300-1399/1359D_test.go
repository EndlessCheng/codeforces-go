package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1359D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5 -2 10 -1 4
outputCopy
6
inputCopy
8
5 2 5 3 -30 -30 6 9
outputCopy
10
inputCopy
3
-10 6 -15
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1359D)
}
