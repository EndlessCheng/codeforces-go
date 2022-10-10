package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1651/D
// https://codeforces.com/problemset/status/1651/problem/D
func TestCF1651D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
2 2
1 2
2 1
3 2
2 3
5 5
outputCopy
1 1
1 1
2 0
3 1
2 4
5 4
inputCopy
8
4 4
2 4
2 2
2 3
1 4
4 2
1 3
3 3
outputCopy
4 3
2 5
2 1
2 5
1 5
4 1
1 2
3 2`
	testutil.AssertEqualCase(t, rawText, 0, CF1651D)
}
