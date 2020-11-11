package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF246C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
1 2 3
outputCopy
1 1
1 2
2 3 2
inputCopy
2 1
7 12
outputCopy
1 12 `
	testutil.AssertEqualCase(t, rawText, 0, CF246C)
}
