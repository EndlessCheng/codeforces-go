package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1490/A
// https://codeforces.com/problemset/status/1490/problem/A
func TestCF1490A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
4
4 2 10 1
2
1 3
2
6 1
3
1 4 2
5
1 2 3 4 3
12
4 31 25 50 30 20 34 46 42 16 15 16
outputCopy
5
1
2
1
0
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1490A)
}
