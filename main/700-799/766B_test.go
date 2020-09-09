package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF766B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 5 3 2 4
outputCopy
YES
inputCopy
3
4 1 2
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF766B)
}
