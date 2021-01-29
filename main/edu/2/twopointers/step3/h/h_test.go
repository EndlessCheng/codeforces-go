package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 7 23 3 5
7 4 3 1 5 8
10 12 7 3 8 9 7
outputCopy
47`
	testutil.AssertEqualCase(t, rawText, 0, run)
}
