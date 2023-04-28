package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1800/problem/G
// https://codeforces.com/problemset/status/1800/problem/G
func TestCF1800G(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
6
1 5
1 6
1 2
2 3
2 4
7
1 5
1 3
3 6
1 4
4 7
4 2
9
1 2
2 4
2 3
3 5
1 7
7 6
7 8
8 9
10
2 9
9 10
2 3
6 7
4 3
1 2
3 8
2 5
6 5
10
3 2
8 10
9 7
4 2
8 2
2 1
4 5
6 5
5 7
1
outputCopy
YES
NO
YES
NO
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1800G)
}
