package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1473/D
// https://codeforces.com/problemset/status/1473/problem/D
func TestCF1473D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
8 4
-+--+--+
1 8
2 8
2 5
1 1
4 10
+-++
1 1
1 2
2 2
1 3
2 3
3 3
1 4
2 4
3 4
4 4
outputCopy
1
2
4
4
3
3
4
2
3
2
1
2
2
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1473D)
}
