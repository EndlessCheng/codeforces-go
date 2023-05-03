package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1778/E
// https://codeforces.com/problemset/status/1778/problem/E
func TestCF1778E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
6
12 12 8 25 6 1
1 5
1 2
2 6
2 3
2 4
3
4 2
3 5
1 2
2
3 8
1 2
4
2 2
2 1
1 2
1 1
3
3 8 7
1 2
2 3
2
2 2
2 1
outputCopy
15
6
29
11
3
8
11
15
3
inputCopy
1
2
3 8
1 2
1
2 2
outputCopy
11`
	testutil.AssertEqualCase(t, rawText, 0, CF1778E)
}
