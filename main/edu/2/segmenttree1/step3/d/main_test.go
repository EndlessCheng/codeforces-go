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
5 1 2 2 3 1 3 4 5 4
outputCopy
1 0 1 1 1 `
	testutil.AssertEqualCase(t, rawText, 0, run)
}
