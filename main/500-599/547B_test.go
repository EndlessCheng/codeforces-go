package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/547/B
// https://codeforces.com/problemset/status/547/problem/B
func TestCF547B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10
1 2 3 4 5 4 3 2 1 6
outputCopy
6 4 4 3 3 2 2 1 1 1 `
	testutil.AssertEqualCase(t, rawText, 0, CF547B)
}
