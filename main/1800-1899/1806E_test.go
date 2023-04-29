package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1806/E
// https://codeforces.com/problemset/status/1806/problem/E
func TestCF1806E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 2
1 5 2 3 1 1
1 2 3 3 2
4 5
6 6
outputCopy
33
27
inputCopy
14 8
3 2 5 3 1 4 2 2 2 5 5 5 2 4
1 2 3 1 1 4 7 3 3 1 5 3 8
4 4
4 10
13 10
3 12
13 9
3 12
9 10
11 5
outputCopy
47
53
48
36
42
36
48
14`
	testutil.AssertEqualCase(t, rawText, 0, CF1806E)
}
