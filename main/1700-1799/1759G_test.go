package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1759/G
// https://codeforces.com/problemset/status/1759/problem/G
func TestCF1759G(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
6
4 3 6
4
2 4
8
8 7 2 3
6
6 4 2
4
4 4
8
8 7 4 5
outputCopy
1 4 2 3 5 6 
1 2 3 4 
-1
5 6 3 4 1 2 
-1
1 8 6 7 2 4 3 5 `
	testutil.AssertEqualCase(t, rawText, 0, CF1759G)
}
