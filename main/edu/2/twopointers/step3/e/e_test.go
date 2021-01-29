package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 20
9 7 6 5 8 4
7 1 3 6 8 3
outputCopy
17`
	testutil.AssertEqualCase(t, rawText, 0, run)
}
