package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/118/E
// https://codeforces.com/problemset/status/118/problem/E
func TestCF118E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 8
1 2
2 3
1 3
4 5
4 6
5 6
2 4
3 5
outputCopy
1 2
2 3
3 1
4 5
5 6
6 4
4 2
3 5
inputCopy
6 7
1 2
2 3
1 3
4 5
4 6
5 6
2 4
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF118E)
}
