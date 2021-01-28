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
4 6 9 3 6
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, run)
}
