package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1672/D
// https://codeforces.com/problemset/status/1672/problem/D
func TestCF1672D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5
1 2 3 3 2
1 3 3 2 2
5
1 2 4 2 1
4 2 2 1 1
5
2 4 5 5 2
2 2 4 5 5
3
1 2 3
1 2 3
3
1 1 2
2 1 1
outputCopy
YES
YES
NO
YES
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1672D)
}
