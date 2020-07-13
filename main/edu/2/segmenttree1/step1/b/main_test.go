package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 5
5 4 2 3 5
2 0 3
1 2 6
2 0 3
1 3 1
2 0 5
outputCopy
2
4
1`
	testutil.AssertEqualCase(t, rawText, 0, run)
}
