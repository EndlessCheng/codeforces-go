package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1328F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 5
1 2 2 4 2 3
outputCopy
3
inputCopy
7 5
3 3 2 1 1 1 3
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1328F)
}
