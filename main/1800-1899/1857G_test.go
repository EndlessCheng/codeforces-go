package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1857/G
// https://codeforces.com/problemset/status/1857/problem/G
func TestCF1857G(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2 5
1 2 4
4 5
1 2 2
2 3 4
3 4 3
5 6
1 2 3
1 3 2
3 4 6
3 5 1
10 200
1 2 3
2 3 33
3 4 200
1 5 132
5 6 1
5 7 29
7 8 187
7 9 20
7 10 4
outputCopy
1
8
80
650867886`
	testutil.AssertEqualCase(t, rawText, 0, CF1857G)
}
