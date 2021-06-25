package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/663/C
// https://codeforces.com/problemset/status/663/problem/C
func TestCF663C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
1 2 B
3 1 R
3 2 B
outputCopy
1
2 
inputCopy
6 5
1 3 R
2 3 R
3 4 B
4 5 R
4 6 R
outputCopy
2
3 4 
inputCopy
4 5
1 2 R
1 3 R
2 3 B
3 4 B
1 4 B
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF663C)
}
