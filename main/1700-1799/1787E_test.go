package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1787/E
// https://codeforces.com/problemset/status/1787/problem/E
func TestCF1787E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
15 6 7
11 4 5
5 3 2
4 1 4
6 1 7
11 5 5
11 6 5
outputCopy
YES
3 6 10 11
3 5 12 14
3 3 9 13
3 1 2 4
2 8 15
1 7
YES
2 1 4
2 2 7
2 3 6
5 5 8 9 10 11
NO
YES
4 1 2 3 4
YES
6 1 2 3 4 5 6
NO
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1787E)
}
