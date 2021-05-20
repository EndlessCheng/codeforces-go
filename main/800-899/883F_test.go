package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/883/F
// https://codeforces.com/problemset/status/883/problem/F
func TestCF883F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10
mihail
oolyana
kooooper
hoon
ulyana
koouper
mikhail
khun
kuooper
kkkhoon
outputCopy
4
inputCopy
9
hariton
hkariton
buoi
kkkhariton
boooi
bui
khariton
boui
boi
outputCopy
5
inputCopy
2
alex
alex
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF883F)
}
