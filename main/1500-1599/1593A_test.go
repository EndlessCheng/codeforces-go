package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1593/A
// https://codeforces.com/problemset/status/1593/problem/A
func TestCF1593A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
0 0 0
10 75 15
13 13 17
1000 0 0
0 1000000000 0
outputCopy
1 1 1
66 0 61
5 5 0
0 1001 1001
1000000001 0 1000000001`
	testutil.AssertEqualCase(t, rawText, 0, CF1593A)
}
