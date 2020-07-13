package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2
5 -4 4 3 -5
4 3
3 -1
outputCopy
8
11
7
inputCopy
4 2
-2 -1 -5 -4
1 3
3 2
outputCopy
0
3
3`
	testutil.AssertEqualCase(t, rawText, 0, run)
}
