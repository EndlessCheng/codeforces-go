package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/453/C
// https://codeforces.com/problemset/status/453/problem/C
func TestCF453C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
1 2
2 3
1 1 1
outputCopy
3
1 2 3
inputCopy
5 7
1 2
1 3
1 4
1 5
3 4
3 5
4 5
0 1 0 1 0
outputCopy
10
2 1 3 4 5 4 5 4 3 1 
inputCopy
2 0
0 0
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF453C)
}
