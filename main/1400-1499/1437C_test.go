package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1437C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
6
4 2 4 4 5 2
7
7 7 7 7 7 7 7
1
1
5
5 1 2 4 3
4
1 4 4 4
21
21 8 1 4 1 5 21 1 8 21 11 21 11 3 12 8 19 15 9 11 13
outputCopy
4
12
0
0
2
21`
	testutil.AssertEqualCase(t, rawText, 0, CF1437C)
}
