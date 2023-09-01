package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1800/problem/D
// https://codeforces.com/problemset/status/1800/problem/D
func TestCF1800D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
6
aaabcc
10
aaaaaaaaaa
6
abcdef
7
abacaba
6
cccfff
4
abba
5
ababa
outputCopy
4
1
5
3
3
3
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1800D)
}
