package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF448C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 2 1 2 1
outputCopy
3
inputCopy
2
2 2
outputCopy
2
inputCopy
1
5
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF448C)
}
