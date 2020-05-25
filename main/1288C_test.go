package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1288C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 2
outputCopy
5
inputCopy
10 1
outputCopy
55
inputCopy
723 9
outputCopy
157557417`
	testutil.AssertEqualCase(t, rawText, 0, CF1288C)
}
