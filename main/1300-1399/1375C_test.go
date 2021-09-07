package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1375/C
// https://codeforces.com/problemset/status/1375/problem/C
func TestCF1375C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3
1 2 3
4
3 1 2 4
3
2 3 1
6
2 4 6 1 3 5
outputCopy
YES
YES
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1375C)
}
