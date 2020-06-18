package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF678D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 4 1 1
outputCopy
7
inputCopy
3 4 2 1
outputCopy
25
inputCopy
3 4 3 1
outputCopy
79`
	testutil.AssertEqualCase(t, rawText, 0, CF678D)
}
