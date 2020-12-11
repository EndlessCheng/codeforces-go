package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1355B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
3
1 1 1
5
2 3 1 2 2
outputCopy
3
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1355B)
}
