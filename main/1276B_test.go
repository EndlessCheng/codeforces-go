package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1276B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
7 7 3 5
1 2
2 3
3 4
4 5
5 6
6 7
7 5
4 5 2 3
1 2
2 3
3 4
4 1
4 2
4 3 2 1
1 2
2 3
4 1
5 5 1 2
1 2
1 4
1 5
2 4
2 3
outputCopy
4
0
1
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1276B)
}
