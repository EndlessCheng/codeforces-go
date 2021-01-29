package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 1
aab
outputCopy
2
inputCopy
6 2
aabcbb
outputCopy
4
inputCopy
0 0
aabbcaabbbbbbbbbbbbbbbb
outputCopy
`
	testutil.AssertEqualCase(t, rawText, 0, run)
}
