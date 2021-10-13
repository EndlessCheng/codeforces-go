package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1515/D
// https://codeforces.com/problemset/status/1515/problem/D
func TestCF1515D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
6 3 3
1 2 3 2 2 2
6 2 4
1 1 2 2 2 2
6 5 1
6 5 4 3 2 1
4 0 4
4 4 4 3
outputCopy
2
3
5
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1515D)
}
