package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1003A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
1 2 4 3 3 2
outputCopy
2
inputCopy
1
100
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1003A)
}
