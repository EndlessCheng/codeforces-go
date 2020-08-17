package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1380D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2
5 2 3
3 1 4 5 2
3 5
outputCopy
8
inputCopy
4 4
5 1 4
4 3 1 2
2 4 3 1
outputCopy
-1
inputCopy
4 4
2 1 11
1 3 2 4
1 3 2 4
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1380D)
}
