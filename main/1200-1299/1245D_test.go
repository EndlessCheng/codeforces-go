package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1245/D
// https://codeforces.com/problemset/status/1245/problem/D
func TestCF1245D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 3
1 1
3 2
3 2 3
3 2 3
outputCopy
8
3
1 2 3 
0
inputCopy
3
2 1
1 2
3 3
23 2 23
3 2 3
outputCopy
27
1
2 
2
1 2
2 3`
	testutil.AssertEqualCase(t, rawText, 0, CF1245D)
}
