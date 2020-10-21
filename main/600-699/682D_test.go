package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF682D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2 2
abc
ab
outputCopy
2
inputCopy
9 12 4
bbaaababb
abbbabbaaaba
outputCopy
7
inputCopy
5 9 1
babcb
abbcbaacb
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, -1, CF682D)
}
