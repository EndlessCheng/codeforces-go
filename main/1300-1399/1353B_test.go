package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1353B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 1
1 2
3 4
5 5
5 5 6 6 5
1 2 5 4 3
5 3
1 2 3 4 5
10 9 10 10 9
4 0
2 2 4 3
2 4 2 3
4 4
1 2 2 1
4 4 5 4
outputCopy
6
27
39
11
17`
	testutil.AssertEqualCase(t, rawText, 0, CF1353B)
}
