package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1471B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 2
12
4 2
4 6 8 2
outputCopy
36
44`
	testutil.AssertEqualCase(t, rawText, 0, CF1471B)
}
