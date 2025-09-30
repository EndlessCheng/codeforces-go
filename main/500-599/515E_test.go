package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/515/E
// https://codeforces.com/problemset/status/515/problem/E
func TestCF515E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 3
2 2 2 2 2
3 5 2 1 4
1 3
2 2
4 5
outputCopy
12
16
18
inputCopy
3 3
5 1 4
5 1 4
3 3
2 2
1 1
outputCopy
17
22
11`
	testutil.AssertEqualCase(t, rawText, 0, CF515E)
}
