package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF627A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9 5
outputCopy
4
inputCopy
3 3
outputCopy
2
inputCopy
5 2
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF627A)
}
