package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/327/E
// https://codeforces.com/problemset/status/327/problem/E
func TestCF327E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 3 5
2
5 7
outputCopy
1
inputCopy
3
2 2 2
2
1 3
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF327E)
}
