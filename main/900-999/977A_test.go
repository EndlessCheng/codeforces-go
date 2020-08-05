package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF977A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
512 4
outputCopy
50
inputCopy
1000000000 9
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF977A)
}
