package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1168A(t *testing.T) {
	// just copy from website
	rawText := `
10 10
5 0 5 9 4 6 4 5 0 0
outputCopy
6
inputCopy
5 3
0 0 0 1 2
outputCopy
0
inputCopy
5 7
0 6 1 3 2
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, -1, CF1168A)
}
