package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1312A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
6 3
7 3
outputCopy
YES
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1312A)
}
