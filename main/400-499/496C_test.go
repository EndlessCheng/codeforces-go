package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/496/C
// https://codeforces.com/problemset/status/496/problem/C
func TestCF496C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 10
codeforces
outputCopy
0
inputCopy
4 4
case
care
test
code
outputCopy
2
inputCopy
5 4
code
forc
esco
defo
rces
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF496C)
}
