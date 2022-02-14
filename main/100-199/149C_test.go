package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/149/C
// https://codeforces.com/problemset/status/149/problem/C
func TestCF149C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2 1
outputCopy
2
1 2 
1
3 
inputCopy
5
2 3 3 1 1
outputCopy
3
4 1 3 
2
5 2 `
	testutil.AssertEqualCase(t, rawText, 0, CF149C)
}
