package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1359A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
8 3 2
4 2 4
9 6 3
42 0 7
outputCopy
3
0
1
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1359A)
}
