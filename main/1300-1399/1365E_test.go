package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1365E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 1 3
outputCopy
3
inputCopy
3
3 1 4
outputCopy
7
inputCopy
1
1
outputCopy
1
inputCopy
4
7 7 1 1
outputCopy
7`
	testutil.AssertEqualCase(t, rawText, 0, CF1365E)
}
