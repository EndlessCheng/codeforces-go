package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF798C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 1
outputCopy
YES
1
inputCopy
3
6 2 4
outputCopy
YES
0
inputCopy
2
1 3
outputCopy
YES
1`
	testutil.AssertEqualCase(t, rawText, 0, CF798C)
}
