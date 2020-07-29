package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF917B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 4
1 2 b
1 3 a
2 4 c
3 4 b
outputCopy
BAAA
ABAA
BBBA
BBBB
inputCopy
5 8
5 3 h
1 2 c
3 1 c
3 2 r
5 1 r
4 3 z
5 4 r
5 2 h
outputCopy
BABBB
BBBBB
AABBB
AAABA
AAAAB`
	testutil.AssertEqualCase(t, rawText, 0, CF917B)
}
