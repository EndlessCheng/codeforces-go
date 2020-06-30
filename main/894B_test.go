package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF894B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1 -1
outputCopy
1
inputCopy
1 3 1
outputCopy
1
inputCopy
3 3 -1
outputCopy
16`
	testutil.AssertEqualCase(t, rawText, 0, CF894B)
}
