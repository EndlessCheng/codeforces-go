package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1585/D
// https://codeforces.com/problemset/status/1585/problem/D
func TestCF1585D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
1
1
2
2 2
2
2 1
3
1 2 3
3
2 1 3
3
3 1 2
4
2 1 4 3
outputCopy
YES
YES
NO
YES
NO
YES
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1585D)
}
