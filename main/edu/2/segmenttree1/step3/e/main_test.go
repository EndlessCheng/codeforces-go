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
1 0 3 3
2 1
1 2 4 4
2 3
2 4
outputCopy
3
4
0`
	testutil.AssertEqualCase(t, rawText, 0, run)
}
