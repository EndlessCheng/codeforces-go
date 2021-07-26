package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/960/F
// https://codeforces.com/problemset/status/960/problem/F
func TestCF960F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
3 1 3
1 2 1
2 3 2
outputCopy
2
inputCopy
5 5
1 3 2
3 2 3
3 4 5
5 4 0
4 5 8
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, -1, CF960F)
}
