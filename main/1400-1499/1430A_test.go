package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1430A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
30
67
4
14
outputCopy
2 2 2
7 5 3
-1
0 0 2`
	testutil.AssertEqualCase(t, rawText, 0, CF1430A)
}
