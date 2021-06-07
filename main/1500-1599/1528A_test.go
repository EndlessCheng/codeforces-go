package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1528/A
// https://codeforces.com/problemset/status/1528/problem/A
func TestCF1528A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2
1 6
3 8
1 2
3
1 3
4 6
7 9
1 2
2 3
6
3 14
12 20
12 19
2 12
10 17
3 17
3 2
6 5
1 5
2 6
4 6
outputCopy
7
8
62`
	testutil.AssertEqualCase(t, rawText, 0, CF1528A)
}
