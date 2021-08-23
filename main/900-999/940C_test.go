package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/940/C
// https://codeforces.com/problemset/status/940/problem/C
func TestCF940C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
abc
outputCopy
aca
inputCopy
3 2
abc
outputCopy
ac
inputCopy
3 3
ayy
outputCopy
yaa
inputCopy
2 3
ba
outputCopy
baa`
	testutil.AssertEqualCase(t, rawText, 0, CF940C)
}
