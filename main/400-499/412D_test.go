package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/412/D
// https://codeforces.com/problemset/status/412/problem/D
func TestCF412D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 1
1 2
outputCopy
2 1 
inputCopy
3 3
1 2
2 3
3 1
outputCopy
2 1 3 `
	testutil.AssertEqualCase(t, rawText, 0, CF412D)
}
