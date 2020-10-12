package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1427D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3 1 2 4
outputCopy
2
3 1 2 1
2 1 3
inputCopy
6
6 5 4 3 2 1
outputCopy
1
6 1 1 1 1 1 1
inputCopy
1
1
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1427D)
}
