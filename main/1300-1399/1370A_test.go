package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1370A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
3
5
outputCopy
1
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1370A)
}
