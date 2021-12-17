package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1609/D
// https://codeforces.com/problemset/status/1609/problem/D
func TestCF1609D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 6
1 2
3 4
2 4
7 6
6 5
1 7
outputCopy
1
1
3
3
3
6
inputCopy
10 8
1 2
2 3
3 4
1 4
6 7
8 9
8 10
1 4
outputCopy
1
2
3
4
5
5
6
8`
	testutil.AssertEqualCase(t, rawText, -1, CF1609D)
}
