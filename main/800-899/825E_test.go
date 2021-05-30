package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/825/E
// https://codeforces.com/problemset/status/825/problem/E
func TestCF825E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
1 2
1 3
3 2
outputCopy
1 3 2 
inputCopy
4 5
3 1
4 1
2 3
3 4
2 4
outputCopy
4 1 2 3 
inputCopy
5 4
3 1
2 1
2 3
4 5
outputCopy
3 1 2 4 5 `
	testutil.AssertEqualCase(t, rawText, 0, CF825E)
}
