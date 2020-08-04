package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF933A(t *testing.T) {
	// just copy from website
	rawText := `
4
1 2 1 2
outputCopy
4
inputCopy
10
1 1 2 2 2 1 1 2 2 1
outputCopy
9
inputCopy
1
2
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF933A)
}
