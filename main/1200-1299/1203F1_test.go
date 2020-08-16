package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1203F1(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 4
4 6
10 -2
8 -1
outputCopy
YES
inputCopy
3 5
4 -5
4 -2
1 3
outputCopy
YES
inputCopy
4 4
5 2
5 -3
2 1
4 -2
outputCopy
YES
inputCopy
3 10
10 0
10 -10
30 0
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1203F1)
}
