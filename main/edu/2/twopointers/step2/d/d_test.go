package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 20
2 6 4 3 6 8 9
outputCopy
9
inputCopy
3 4
1 1 1
outputCopy
`
	testutil.AssertEqualCase(t, rawText, 0, run)
}
