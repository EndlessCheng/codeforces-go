package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/732/F
// https://codeforces.com/problemset/status/732/problem/F
func TestCF732F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 9
4 3
2 6
7 1
4 1
7 3
3 5
7 4
6 5
2 5
outputCopy
4
4 3
6 2
7 1
1 4
3 7
5 3
7 4
5 6
2 5`
	testutil.AssertEqualCase(t, rawText, 0, CF732F)
}
