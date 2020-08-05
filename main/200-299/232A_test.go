package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF232A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
outputCopy
3
011
101
110
inputCopy
10
outputCopy
5
01111
10111
11011
11101
11110`
	testutil.AssertEqualCase(t, rawText, 0, CF232A)
}
