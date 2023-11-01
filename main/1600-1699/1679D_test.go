package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1679/D
// https://codeforces.com/problemset/status/1679/problem/D
func TestCF1679D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 7 4
1 10 2 3 4 5
1 2
1 3
3 4
4 5
5 6
6 2
2 5
outputCopy
4
inputCopy
6 7 100
1 10 2 3 4 5
1 2
1 3
3 4
4 5
5 6
6 2
2 5
outputCopy
10
inputCopy
2 1 5
1 1
1 2
outputCopy
-1
inputCopy
1 0 1
1000000000
outputCopy
1000000000`
	testutil.AssertEqualCase(t, rawText, 0, CF1679D)
}
