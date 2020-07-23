package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF474D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
1 3
2 3
4 4
outputCopy
6
5
5`
	testutil.AssertEqualCase(t, rawText, 0, CF474D)
}
