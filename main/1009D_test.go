package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1009D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 6
outputCopy
Possible
2 5
3 2
5 1
3 4
4 1
5 4
inputCopy
6 12
outputCopy
Impossible`
	testutil.AssertEqualCase(t, rawText, 0, CF1009D)
}
