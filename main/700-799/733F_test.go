package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/733/F
// https://codeforces.com/problemset/status/733/problem/F
func TestCF733F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 9
1 3 1 1 3 1 2 2 2
4 1 4 2 2 5 3 1 6
1 2
1 3
2 3
2 4
2 5
3 5
3 6
4 5
5 6
7
outputCopy
0
1 1
3 1
6 1
7 2
8 -5
inputCopy
3 3
9 5 1
7 7 2
2 1
3 1
3 2
2
outputCopy
5
3 0
2 5`
	testutil.AssertEqualCase(t, rawText, 0, CF733F)
}
