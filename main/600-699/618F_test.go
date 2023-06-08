package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/618/F
// https://codeforces.com/problemset/status/618/problem/F
func TestCF618F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10
10 10 10 10 10 10 10 10 10 10
10 9 8 7 6 5 4 3 2 1
outputCopy
1
2
3
5 8 10
inputCopy
5
4 4 3 3 3
2 2 2 2 5
outputCopy
2
2 3
2
3 5`
	testutil.AssertEqualCase(t, rawText, 0, CF618F)
}
