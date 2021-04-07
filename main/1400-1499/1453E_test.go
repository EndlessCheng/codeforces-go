package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1453/E
// https://codeforces.com/problemset/status/1453/problem/E
func TestCF1453E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3
1 2
1 3
4
1 2
2 3
3 4
8
1 2
2 3
3 4
1 5
5 6
6 7
5 8
outputCopy
2
3
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1453E)
}
