package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF808D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 3 2
outputCopy
YES
inputCopy
5
1 2 3 4 5
outputCopy
NO
inputCopy
5
2 2 3 4 5
outputCopy
YES
inputCopy
6
6 100 100 3 4 1
outputCopy
YES
inputCopy
4
6 1 4 5
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, -1, CF808D)
}
