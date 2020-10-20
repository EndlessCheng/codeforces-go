package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1223D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
7
3 1 6 6 3 1 1
8
1 1 4 4 4 7 8 8
7
4 2 5 2 6 2 7
outputCopy
2
0
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1223D)
}
