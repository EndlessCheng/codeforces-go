package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF894C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2 4 6 12
outputCopy
3
4 6 12
inputCopy
2
2 3
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF894C)
}
