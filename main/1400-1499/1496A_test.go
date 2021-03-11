package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1496/problem/A
// https://codeforces.com/problemset/status/1496/problem/A
func TestCF1496A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
5 1
qwqwq
2 1
ab
3 1
ioi
4 2
icpc
22 0
dokidokiliteratureclub
19 8
imteamshanghaialice
6 3
aaaaaa
outputCopy
YES
NO
YES
NO
YES
NO
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1496A)
}
