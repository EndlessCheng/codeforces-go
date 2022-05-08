package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1675/E
// https://codeforces.com/problemset/status/1675/problem/E
func TestCF1675E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3 2
cba
4 5
fgde
7 5
gndcafb
4 19
ekyv
outputCopy
aaa
agaa
bnbbabb
aapp
inputCopy
1
5 19
lssls
outputCopy
aaaaa`
	testutil.AssertEqualCase(t, rawText, 0, CF1675E)
}
