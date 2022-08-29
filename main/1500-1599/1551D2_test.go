package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1551/D2
// https://codeforces.com/problemset/status/1551/problem/D2
func TestCF1551D2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
4 4 2
2 3 0
3 2 3
1 2 0
2 4 2
5 2 2
2 17 16
2 1 1
outputCopy
YES
accx
aegx
bega
bdda
YES
aha
aha
YES
zz
aa
zz
NO
YES
aaza
bbza
NO
YES
bbaabbaabbaabbaay
ddccddccddccddccy
NO
inputCopy
1
2 2 1
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1551D2)
}
