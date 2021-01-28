package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 7
1 6 9 13 18 18
2 3 8 13 15 21 25
outputCopy
1 2 3 6 8 9 13 13 15 18 18 21 25`
	testutil.AssertEqualCase(t, rawText, 0, run)
}
