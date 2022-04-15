package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1661/A
// https://codeforces.com/problemset/status/1661/problem/A
func TestCF1661A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4
3 3 10 10
10 10 3 3
5
1 2 3 4 5
6 7 8 9 10
6
72 101 108 108 111 44
10 87 111 114 108 100
outputCopy
0
8
218`
	testutil.AssertEqualCase(t, rawText, 0, CF1661A)
}
