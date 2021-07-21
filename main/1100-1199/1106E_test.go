package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1106/E
// https://codeforces.com/problemset/status/1106/problem/E
func TestCF1106E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 0 2
1 3 4 5
2 5 5 8
outputCopy
13
inputCopy
10 1 6
1 1 2 4
2 2 6 2
3 3 3 3
4 4 4 5
5 5 5 7
6 6 6 9
outputCopy
2
inputCopy
12 2 6
1 5 5 4
4 6 6 2
3 8 8 3
2 9 9 5
6 10 10 7
8 12 12 9
outputCopy
11`
	testutil.AssertEqualCase(t, rawText, 0, CF1106E)
}
