package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/723/E
// https://codeforces.com/problemset/status/723/problem/E
func TestCF723E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
5 5
2 1
4 5
2 3
1 3
3 5
7 2
3 7
4 2
outputCopy
3
1 3
3 5
5 4
3 2
2 1
3
2 4
3 7`
	testutil.AssertEqualCase(t, rawText, 0, CF723E)
}
