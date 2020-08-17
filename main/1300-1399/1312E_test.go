package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1312E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
4 3 2 2 3
outputCopy
2
inputCopy
7
3 3 4 4 4 3 3
outputCopy
2
inputCopy
3
1 3 5
outputCopy
3
inputCopy
1
1000
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1312E)
}
