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
1 3 2 4 3
2 3 0
2 3 2
1 2 5
2 4 1
2 5 4
1 3 7
2 6 1
outputCopy
1
3
2
-1
3`
	testutil.AssertEqualCase(t, rawText, 0, run)
}
