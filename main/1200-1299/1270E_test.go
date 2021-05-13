package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1270/E
// https://codeforces.com/problemset/status/1270/problem/E
func TestCF1270E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
0 0
0 1
1 0
outputCopy
1
1 
inputCopy
4
0 1
0 -1
1 0
-1 0
outputCopy
2
1 2 
inputCopy
3
-2 1
1 1
-1 0
outputCopy
1
2 
inputCopy
6
2 5
0 3
-4 -1
-5 -4
1 0
3 -1
outputCopy
1
6 
inputCopy
2
-1000000 -1000000
1000000 1000000
outputCopy
1
1 
inputCopy
4
0 0
2 0
0 2
2 2
outputCopy
2
1 4 `
	testutil.AssertEqualCase(t, rawText, 0, CF1270E)
}
