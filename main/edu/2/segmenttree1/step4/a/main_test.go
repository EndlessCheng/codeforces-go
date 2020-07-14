package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2 3
5
1 1 2
1 1 3
1 2 3
0 2 1
1 1 3
outputCopy
-1
2
-1
3`
	testutil.AssertEqualCase(t, rawText, 0, run)
}
