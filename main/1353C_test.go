package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1353C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1
5
499993
outputCopy
0
40
41664916690999888`
	testutil.AssertEqualCase(t, rawText, 0, CF1353C)
}
