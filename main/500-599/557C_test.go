package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF557C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 5
3 2
outputCopy
2
inputCopy
3
2 4 4
1 1 1
outputCopy
0
inputCopy
6
2 2 1 1 3 3
4 3 5 5 2 1
outputCopy
8`
	testutil.AssertEqualCase(t, rawText, 0, CF557C)
}
