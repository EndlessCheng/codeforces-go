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
4 1 3 5 2
outputCopy
0 1 1 0 3 `
	testutil.AssertEqualCase(t, rawText, 0, run)
}
