package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF633D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2 -1
outputCopy
3
inputCopy
5
28 35 7 14 21
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF633D)
}