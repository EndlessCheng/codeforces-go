package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1370D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 2
1 2 3 4
outputCopy
1
inputCopy
4 3
1 2 3 4
outputCopy
2
inputCopy
5 3
5 3 4 2 6
outputCopy
2
inputCopy
6 4
5 3 50 2 4 5
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1370D)
}
