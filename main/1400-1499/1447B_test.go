package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1447B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
2 2
-1 1
1 1
3 4
0 -1 -2 -3
-1 -2 -3 -4
-2 -3 -4 -5
outputCopy
2
30`
	testutil.AssertEqualCase(t, rawText, 0, CF1447B)
}
