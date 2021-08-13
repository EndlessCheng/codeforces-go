package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/508/E
// https://codeforces.com/problemset/status/508/problem/E
func TestCF508E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 1
1 1
1 1
1 1
outputCopy
()()()()
inputCopy
3
5 5
3 3
1 1
outputCopy
((()))
inputCopy
3
5 5
3 3
2 2
outputCopy
IMPOSSIBLE
inputCopy
3
2 3
1 4
1 4
outputCopy
(())()`
	testutil.AssertEqualCase(t, rawText, 0, CF508E)
}
