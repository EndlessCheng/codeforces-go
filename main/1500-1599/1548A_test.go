package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1548/A
// https://codeforces.com/problemset/status/1548/problem/A
func TestCF1548A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 3
2 1
1 3
3 4
4
3
1 2 3
2 3 1
3
outputCopy
2
1
inputCopy
4 3
2 3
3 4
4 1
1
3
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1548A)
}
