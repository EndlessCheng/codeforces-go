package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1420/C2
// https://codeforces.com/problemset/status/1420/problem/C2
func TestCF1420C2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3 1
1 3 2
1 2
2 2
1 2
1 2
1 2
7 5
1 2 5 4 3 6 7
1 2
6 7
3 4
1 2
2 3
outputCopy
3
4
2
2
2
9
10
10
10
9
11`
	testutil.AssertEqualCase(t, rawText, 0, CF1420C2)
}
