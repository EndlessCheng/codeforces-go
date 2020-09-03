package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1398B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
01111001
0000
111111
101010101
011011110111
outputCopy
4
0
6
3
6`
	testutil.AssertEqualCase(t, rawText, 0, CF1398B)
}
