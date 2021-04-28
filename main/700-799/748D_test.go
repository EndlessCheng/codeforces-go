package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/748/D
// https://codeforces.com/problemset/status/748/problem/D
func TestCF748D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 3
abb 2
aaa -3
bba -1
zyz -4
abb 5
aaa 7
xyx 4
outputCopy
12
inputCopy
3 1
a 1
a 2
a 3
outputCopy
6
inputCopy
2 5
abcde 10000
abcde 10000
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF748D)
}
