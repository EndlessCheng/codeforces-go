package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1368C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
outputCopy
12
1 0
2 0
0 1
1 1
2 1
3 1
0 2
1 2
2 2
3 2
1 3
2 3`
	testutil.AssertEqualCase(t, rawText, 0, CF1368C)
}
