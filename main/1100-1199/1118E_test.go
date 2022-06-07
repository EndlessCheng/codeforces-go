package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1118/E
// https://codeforces.com/problemset/status/1118/problem/E
func TestCF1118E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 3
outputCopy
YES
3 1
1 3
3 2
2 3
inputCopy
10 4
outputCopy
YES
2 1
1 3
4 2
3 4
4 3
3 2
2 4
4 1
1 4
3 1
inputCopy
13 4
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1118E)
}
