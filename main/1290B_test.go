package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1290B(t *testing.T) {
	// just copy from website
	rawText := `
aaaaa
3
1 1
2 4
5 5
outputCopy
Yes
No
Yes
inputCopy
aabbbbbbc
6
1 2
2 4
2 2
1 9
5 7
3 5
outputCopy
No
Yes
Yes
Yes
No
No`
	testutil.AssertEqualCase(t, rawText, 0, CF1290B)
}
