package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/723/F
// https://codeforces.com/problemset/status/723/problem/F
func TestCF723F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
1 2
2 3
3 1
1 2 1 1
outputCopy
Yes
3 2
1 3
inputCopy
7 8
7 4
1 3
5 4
5 7
3 2
2 4
6 1
1 2
6 4 1 4
outputCopy
Yes
1 3
5 7
3 2
7 4
2 4
6 1
inputCopy
5 5
1 3
1 4
1 5
2 3
2 4
1 2 2 2
outputCopy
Yes
5 1
3 2
4 2
1 3
inputCopy
5 5
1 3
2 4
1 5
2 3
1 4
1 2 2 2
outputCopy
Yes
5 1
3 2
4 2
1 3`
	testutil.AssertEqualCase(t, rawText, 2, CF723F)
}
