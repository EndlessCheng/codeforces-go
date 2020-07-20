package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF978G(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2
1 3 1
1 5 1
outputCopy
1 2 3 0 3 
inputCopy
3 2
1 3 1
1 2 1
outputCopy
-1
inputCopy
10 3
4 7 2
1 10 3
8 9 1
outputCopy
2 2 2 1 1 0 4 3 4 4 `
	testutil.AssertEqualCase(t, rawText, 0, CF978G)
}
