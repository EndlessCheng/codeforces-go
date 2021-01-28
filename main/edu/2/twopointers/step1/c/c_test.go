package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8 7
1 1 3 3 3 5 8 8
1 3 3 4 5 5 5
outputCopy
11`
	testutil.AssertEqualCase(t, rawText, 0, run)
}
