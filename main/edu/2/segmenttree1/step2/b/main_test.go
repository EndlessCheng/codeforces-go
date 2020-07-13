package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 7
1 1 0 1 0
2 0
2 1
2 2
1 2
2 3
1 0
2 0
outputCopy
0
1
3
3
1`
	testutil.AssertEqualCase(t, rawText, 0, run)
}
