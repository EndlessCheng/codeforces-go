package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1619/D
// https://codeforces.com/problemset/status/1619/problem/D
func TestCF1619D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5

2 2
1 2
3 4

4 3
1 3 1
3 1 1
1 2 2
1 1 3

2 3
5 3 4
2 5 1

4 2
7 9
8 1
9 6
10 8

2 4
6 5 2 1
7 9 7 2
outputCopy
3
2
4
8
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1619D)
}
