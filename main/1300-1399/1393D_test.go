package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1393/problem/D
// https://codeforces.com/problemset/status/1393/problem/D
func TestCF1393D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
aaa
aaa
aaa
outputCopy
10
inputCopy
3 4
abab
baba
abab
outputCopy
12
inputCopy
5 5
zbacg
baaac
aaaaa
eaaad
weadd
outputCopy
31`
	testutil.AssertEqualCase(t, rawText, 0, CF1393D)
}
