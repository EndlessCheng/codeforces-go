package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/723/C
// https://codeforces.com/problemset/status/723/problem/C
func TestCF723C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 2
1 2 3 2
outputCopy
2 1
1 2 1 2
inputCopy
7 3
1 3 2 2 2 2 1
outputCopy
2 1
1 3 3 2 2 2 1
inputCopy
4 4
1000000000 100 7 1000000000
outputCopy
1 4
1 2 3 4 
inputCopy
5 4
3 1 495987801 522279660 762868488
outputCopy
1 2
3 1 2 4 762868488 `
	testutil.AssertEqualCase(t, rawText, 0, CF723C)
}
