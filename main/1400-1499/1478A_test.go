package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1478/A
// https://codeforces.com/problemset/status/1478/problem/A
func TestCF1478A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
6
1 1 1 2 3 4
5
1 1 2 2 3
4
2 2 2 2
3
1 2 3
1
1
outputCopy
3
2
4
1
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1478A)
}
