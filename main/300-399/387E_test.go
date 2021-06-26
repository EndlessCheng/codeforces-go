package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/387/E
// https://codeforces.com/problemset/status/387/problem/E
func TestCF387E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
2 1 3
1 3
outputCopy
1
inputCopy
10 5
1 2 3 4 5 6 7 8 9 10
2 4 6 8 10
outputCopy
30`
	testutil.AssertEqualCase(t, rawText, 0, CF387E)
}
