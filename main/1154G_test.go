package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1154G(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 4 8 3 6
outputCopy
1 2
inputCopy
5
5 2 11 3 7
outputCopy
2 4
inputCopy
6
2 5 10 1 10 2
outputCopy
1 4`
	testutil.AssertEqualCase(t, rawText, 0, CF1154G)
}
