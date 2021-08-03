package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/983/B
// https://codeforces.com/problemset/status/983/problem/B
func TestCF983B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
8 4 1
2
2 3
1 2
outputCopy
5
12
inputCopy
6
1 2 4 8 16 32
4
1 6
2 5
3 4
1 2
outputCopy
60
30
12
3`
	testutil.AssertEqualCase(t, rawText, 0, CF983B)
}
