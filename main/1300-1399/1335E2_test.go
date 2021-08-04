package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1335/E2
// https://codeforces.com/problemset/status/1335/problem/E2
func TestCF1335E2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
8
1 1 2 2 3 2 1 1
3
1 3 3
4
1 10 10 1
1
26
2
2 1
3
1 1 1
outputCopy
7
2
4
1
1
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1335E2)
}
