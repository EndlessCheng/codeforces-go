package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1409E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
7 1
1 5 2 3 1 5 4
1 3 6 7 2 5 4
1 1
1000000000
1000000000
5 10
10 7 5 15 8
20 199 192 219 1904
10 10
15 19 8 17 20 10 9 2 10 19
12 13 6 17 1 14 7 9 19 3
outputCopy
6
1
5
10`
	testutil.AssertEqualCase(t, rawText, 0, CF1409E)
}
