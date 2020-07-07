package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1009E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 2
outputCopy
5
inputCopy
4
1 3 3 7
outputCopy
60`
	testutil.AssertEqualCase(t, rawText, 0, CF1009E)
}
