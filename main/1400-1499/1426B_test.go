package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1426B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
3 4
1 2
5 6
5 7
7 4
8 9
9 8
2 5
1 1
1 1
2 2
2 2
1 100
10 10
10 10
1 2
4 5
8 4
2 2
1 1
1 1
1 2
3 4
1 2
1 1
1 1
outputCopy
YES
NO
YES
NO
YES
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1426B)
}
