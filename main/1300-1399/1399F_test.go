package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1399/problem/F
// https://codeforces.com/problemset/status/1399/problem/F
func TestCF1399F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4
1 5
2 4
2 3
3 4
5
1 5
2 3
2 5
3 5
2 2
3
1 3
2 4
2 3
7
1 10
2 8
2 5
3 4
4 4
6 8
7 7
outputCopy
3
4
2
7`
	testutil.AssertEqualCase(t, rawText, 0, CF1399F)
}
