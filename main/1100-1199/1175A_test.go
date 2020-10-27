package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1175A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
59 3
1000000000000000000 10
outputCopy
8
19`
	testutil.AssertEqualCase(t, rawText, 0, CF1175A)
}
