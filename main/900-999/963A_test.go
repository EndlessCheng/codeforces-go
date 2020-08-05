package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF963A(t *testing.T) {
	// just copy from website
	rawText := `
2 2 3 3
+-+
outputCopy
7
inputCopy
4 1 5 1
-
outputCopy
999999228
inputCopy
1 1 4 2
-+
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 3, CF963A)
}
