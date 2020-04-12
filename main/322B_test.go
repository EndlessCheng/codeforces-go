package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF322B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 6 9
outputCopy
6
inputCopy
4 4 4
outputCopy
4
inputCopy
0 0 0
outputCopy
0
inputCopy
0 2 2
outputCopy
0
inputCopy
3 2 2
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF322B)
}
