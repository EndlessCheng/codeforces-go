package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF977D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
4 8 6 3 12 9
outputCopy
9 3 6 12 4 8 
inputCopy
4
42 28 84 126
outputCopy
126 42 84 28 
inputCopy
2
1000000000000000000 3000000000000000000
outputCopy
3000000000000000000 1000000000000000000 `
	testutil.AssertEqualCase(t, rawText, 0, CF977D)
}
