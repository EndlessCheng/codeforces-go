package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/883/E
// https://codeforces.com/problemset/status/883/problem/E
func TestCF883E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
a**d
2
abcd
acbd
outputCopy
2
inputCopy
5
lo*er
2
lover
loser
outputCopy
0
inputCopy
3
a*a
2
aaa
aba
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF883E)
}
