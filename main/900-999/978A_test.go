package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF978A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
1 5 5 1 6 1
outputCopy
3
5 6 1 
inputCopy
5
2 4 2 4 4
outputCopy
2
2 4 
inputCopy
5
6 6 6 6 6
outputCopy
1
6 `
	testutil.AssertEqualCase(t, rawText, 0, CF978A)
}
