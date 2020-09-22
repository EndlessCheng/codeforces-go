package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1076D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3 2
1 2 1
3 2 1
1 3 3
outputCopy
2
1 2 
inputCopy
4 5 2
4 1 8
2 4 1
2 1 3
3 4 9
3 1 5
outputCopy
2
3 2 `
	testutil.AssertEqualCase(t, rawText, 0, CF1076D)
}
