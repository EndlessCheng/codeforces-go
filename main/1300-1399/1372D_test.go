package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1372D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
7 10 2
outputCopy
17
inputCopy
1
4
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1372D)
}
