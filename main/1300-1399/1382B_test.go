package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1382/B
// https://codeforces.com/problemset/status/1382/problem/B
func TestCF1382B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
3
2 5 4
8
1 1 1 1 1 1 1 1
6
1 2 3 4 5 6
6
1 1 2 1 2 2
1
1000000000
5
1 2 2 1 1
3
1 1 1
outputCopy
First
Second
Second
First
First
Second
First`
	testutil.AssertEqualCase(t, rawText, 0, CF1382B)
}
