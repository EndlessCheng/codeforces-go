package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1610/E
// https://codeforces.com/problemset/status/1610/problem/E
func TestCF1610E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3
1 2 3
5
1 4 4 5 6
6
7 8 197860736 212611869 360417095 837913434
8
6 10 56026534 405137099 550504063 784959015 802926648 967281024
outputCopy
0
1
2
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1610E)
}
