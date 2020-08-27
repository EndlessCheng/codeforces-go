package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF13C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3 2 -1 2 11
outputCopy
4
inputCopy
5
2 1 1 1 1
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF13C)
}
