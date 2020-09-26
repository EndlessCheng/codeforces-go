package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF543B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 4
1 2
2 3
3 4
4 5
1 3 2
3 5 2
outputCopy
0
inputCopy
5 4
1 2
2 3
3 4
4 5
1 3 2
2 4 2
outputCopy
1
inputCopy
5 4
1 2
2 3
3 4
4 5
1 3 2
3 5 1
outputCopy
-1
inputCopy
10 11
1 3
2 3
3 4
4 5
4 6
3 7
3 8
4 9
4 10
7 9
8 10
1 5 3
6 2 3
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF543B)
}
