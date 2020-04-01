package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1081D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 3 2
2 1
1 2 3
1 2 2
2 2 1
outputCopy
2 2 
inputCopy
4 5 3
1 2 3
1 2 5
4 2 1
2 3 2
1 4 4
1 3 3
outputCopy
3 3 3 `
	testutil.AssertEqualCase(t, rawText, 0, CF1081D)
}
