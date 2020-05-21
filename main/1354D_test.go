package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1354D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 5
1 2 3 4 5
-1 -1 -1 -1 -1
outputCopy
0
inputCopy
5 4
1 2 3 4 5
-5 -1 -3 -1
outputCopy
3
inputCopy
6 2
1 1 1 2 3 4
5 6
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF1354D)
}
