package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8 3
1 2 1 2 1 1 2 1
outputCopy
9
3
1 2 1 
inputCopy
1 1
1
outputCopy
1
1
1 
inputCopy
6 2
1 2 1 2 1 2
outputCopy
8
4
1 2 1 2 `
	testutil.AssertEqualCase(t, rawText, 0, run)
}
