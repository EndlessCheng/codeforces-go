package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1119E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2 2 2 2
outputCopy
3
inputCopy
3
1 1 1
outputCopy
0
inputCopy
3
3 3 3
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1119E)
}
