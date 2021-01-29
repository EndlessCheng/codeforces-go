package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 4
1 3 5 8
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, run)
}
