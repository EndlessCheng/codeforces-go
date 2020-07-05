package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1375D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3
2 2 3
3
2 1 0
7
0 7 3 1 3 7 7
9
2 0 1 1 2 4 4 2 0
9
8 4 7 6 1 2 3 0 5
outputCopy
0

2
3 1
4
2 5 5 4
11
3 8 9 7 8 5 9 6 4 1 2
10
1 8 1 9 5 2 4 6 3 7`
	testutil.AssertEqualCase(t, rawText, 0, CF1375D)
}
