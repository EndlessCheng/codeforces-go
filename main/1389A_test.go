package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1389A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 1337
13 69
2 4
88 89
outputCopy
6 7
14 21
2 4
-1 -1`
	testutil.AssertEqualCase(t, rawText, 0, CF1389A)
}
