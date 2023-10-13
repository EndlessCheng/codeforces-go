package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1765/N
// https://codeforces.com/problemset/status/1765/problem/N
func TestCF1765N(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
10000
4
1337
0
987654321
6
66837494128
5
7808652
3
outputCopy
1
1337
321
344128
7052
inputCopy
1
2010
2
outputCopy
10`
	testutil.AssertEqualCase(t, rawText, 0, CF1765N)
}
