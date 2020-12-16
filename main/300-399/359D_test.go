package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF359D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
4 6 9 3 6
outputCopy
1 3
2 
inputCopy
5
1 3 5 7 9
outputCopy
1 4
1 
inputCopy
5
2 3 5 7 11
outputCopy
5 0
1 2 3 4 5 `
	testutil.AssertEqualCase(t, rawText, 0, CF359D)
}
