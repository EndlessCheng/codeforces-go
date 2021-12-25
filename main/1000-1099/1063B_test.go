package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1063/B
// https://codeforces.com/problemset/status/1063/problem/B
func TestCF1063B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 5
3 2
1 2
.....
.***.
...**
*....
outputCopy
10
inputCopy
4 4
2 2
0 1
....
..*.
....
....
outputCopy
7`
	testutil.AssertEqualCase(t, rawText, 0, CF1063B)
}
