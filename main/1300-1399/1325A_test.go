package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1325A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
2
14
outputCopy
1 1
6 4`
	testutil.AssertEqualCase(t, rawText, 0, CF1325A)
}
