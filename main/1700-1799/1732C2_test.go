package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1732/C2
// https://codeforces.com/problemset/status/1732/problem/C2
func TestCF1732C2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
1 1
0
1 1
2 2
5 10
1 2
2 2
3 3
0 2 4
1 3
1 2
2 3
4 4
0 12 8 3
1 4
1 3
2 4
2 3
5 5
21 32 32 32 10
1 5
1 4
1 3
2 5
3 5
7 7
0 1 0 1 0 1 0
1 7
3 6
2 5
1 4
4 7
2 6
2 7
outputCopy
1 1
1 1
2 2
1 1
1 1
2 2
2 3
2 3
2 3
2 3
2 3
2 3
2 3
2 3
3 4
2 4
4 6
2 4
2 4
4 6
2 4
2 4`
	testutil.AssertEqualCase(t, rawText, 0, CF1732C2)
}
