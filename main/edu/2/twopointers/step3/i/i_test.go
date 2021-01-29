package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10 100
14 33 22 21 11 5 13 28 61 2
outputCopy
5
inputCopy
2 2
1 1
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, run)
}
