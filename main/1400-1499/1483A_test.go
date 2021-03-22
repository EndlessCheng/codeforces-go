package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1483/A
// https://codeforces.com/problemset/status/1483/problem/A
func TestCF1483A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
4 6
1 1
2 1 2
3 1 2 3
4 1 2 3 4
2 2 3
1 3
2 2
1 1
1 1
outputCopy
YES
1 2 1 1 2 3 
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1483A)
}
