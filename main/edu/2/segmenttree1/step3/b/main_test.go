package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
0 1 1 0 3
outputCopy
4 1 3 5 2
inputCopy
1
0
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, run)
}
