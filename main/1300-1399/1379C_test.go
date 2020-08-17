package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1379C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
4 3
5 0
1 4
2 2

5 3
5 2
4 2
3 1
outputCopy
14
16`
	testutil.AssertEqualCase(t, rawText, 0, CF1379C)
}
