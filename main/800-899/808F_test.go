package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF808F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 8
5 5 1
1 5 4
4 6 3
1 12 4
3 12 1
outputCopy
4
inputCopy
3 7
4 4 1
5 8 2
5 3 3
outputCopy
2
inputCopy
10 20
9 4 10
2 8 9
9 1 1
8 10 10
5 10 2
1 2 10
9 6 3
2 10 5
7 10 6
6 3 1
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF808F)
}
