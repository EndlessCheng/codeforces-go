package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1562/problem/E
// https://codeforces.com/problemset/status/1562/problem/E
func TestCF1562E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
5
acbac
8
acabacba
12
aaaaaaaaaaaa
10
abacabadac
8
dcbaabcd
3
cba
6
sparky
outputCopy
9
17
12
29
14
3
9`
	testutil.AssertEqualCase(t, rawText, 0, CF1562E)
}
