package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1280/B
// https://codeforces.com/problemset/status/1280/problem/B
func TestCF1280B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
7 8
AAPAAAAA
PPPPAAAA
PPPPAAAA
APAAPPPP
APAPPAPP
AAAAPPAP
AAAAPPAA
6 5
AAAAA
AAAAA
AAPAA
AAPAP
AAAPP
AAAPP
4 4
PPPP
PPPP
PPPP
PPPP
3 4
PPPP
PAAP
PPPP
outputCopy
2
1
MORTAL
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1280B)
}
