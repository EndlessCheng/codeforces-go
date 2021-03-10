package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF10C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
outputCopy
2
inputCopy
5
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF10C)
}
