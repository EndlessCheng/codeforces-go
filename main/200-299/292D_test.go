package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/292/D
// https://codeforces.com/problemset/status/292/problem/D
func TestCF292D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 5
1 2
5 4
2 3
3 1
3 6
6
1 3
2 5
1 5
5 5
2 4
3 3
outputCopy
4
5
6
3
4
2`
	testutil.AssertEqualCase(t, rawText, 0, CF292D)
}
