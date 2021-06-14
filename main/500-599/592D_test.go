package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/592/D
// https://codeforces.com/problemset/status/592/problem/D
func TestCF592D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 2
1 2
1 3
1 4
3 5
3 6
3 7
2 7
outputCopy
2
3
inputCopy
6 4
1 2
2 3
2 4
4 5
4 6
2 4 5 6
outputCopy
2
4`
	testutil.AssertEqualCase(t, rawText, 0, CF592D)
}
