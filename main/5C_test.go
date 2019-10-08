package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol5C(t *testing.T) {
	// just copy from website
	rawText := `
((((()(((
outputCopy
2 1
inputCopy
)((())))(()())
outputCopy
6 2
inputCopy
))(
outputCopy
0 1`
	testutil.AssertEqualCase(t, rawText, 2, Sol5C)
}
