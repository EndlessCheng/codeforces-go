package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1422C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
107
outputCopy
42
inputCopy
100500100500
outputCopy
428101984`
	testutil.AssertEqualCase(t, rawText, 0, CF1422C)
}
