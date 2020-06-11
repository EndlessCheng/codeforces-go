package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1000E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 5
1 2
2 3
3 1
4 1
5 2
outputCopy
2
inputCopy
4 3
1 2
4 3
3 2
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1000E)
}
