package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1675/F
// https://codeforces.com/problemset/status/1675/problem/F
func TestCF1675F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3

3 1
1 3
2
1 3
1 2

6 4
3 5
1 6 2 1
1 3
3 4
3 5
5 6
5 2

6 2
3 2
5 3
1 3
3 4
3 5
5 6
5 2
outputCopy
3
7
2
inputCopy
1
4 2
3 2
4 2
1 2
4 3
3 2
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, -1, CF1675F)
}
