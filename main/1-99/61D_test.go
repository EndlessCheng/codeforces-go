package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF61D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2 3
2 3 4
outputCopy
7
inputCopy
3
1 2 3
1 3 3
outputCopy
9`
	testutil.AssertEqualCase(t, rawText, 0, CF61D)
}
