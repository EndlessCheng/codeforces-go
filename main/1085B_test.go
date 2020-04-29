package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1085B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 3
outputCopy
11
inputCopy
1 2
outputCopy
3
inputCopy
4 6
outputCopy
10`
	testutil.AssertEqualCase(t, rawText, 0, CF1085B)
}
