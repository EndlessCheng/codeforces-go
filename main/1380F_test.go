package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1380F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 3
14
2 4
2 1
1 0
outputCopy
15
12
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1380F)
}
