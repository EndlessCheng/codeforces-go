package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1635/E
// https://codeforces.com/problemset/status/1635/problem/E
func TestCF1635E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 4
1 1 2
1 2 3
2 3 4
2 4 1
outputCopy
YES
R 0
L -3
R 5
L 6
inputCopy
3 3
1 1 2
1 2 3
1 1 3
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1635E)
}
