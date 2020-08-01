package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF778A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
ababcba
abb
5 3 4 1 7 6 2
outputCopy
3
inputCopy
bbbabb
bb
1 6 3 4 2 5
outputCopy
4
inputCopy
bbb
bbb
1 2 3
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF778A)
}
