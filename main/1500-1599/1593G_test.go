package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1593/G
// https://codeforces.com/problemset/status/1593/problem/G
func TestCF1593G(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
([))[)()][]]
3
1 12
4 9
3 6
))))))
2
2 3
1 4
[]
1
1 2
outputCopy
0
2
1
0
0
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1593G)
}
