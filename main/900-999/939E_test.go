package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/939/E
// https://codeforces.com/problemset/status/939/problem/E
func TestCF939E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
1 3
2
1 4
2
1 8
2
outputCopy
0.0000000000
0.5000000000
3.0000000000
inputCopy
4
1 1
1 4
1 5
2
outputCopy
2.0000000000`
	testutil.AssertEqualCase(t, rawText, 0, CF939E)
}
