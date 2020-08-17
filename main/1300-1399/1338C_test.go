package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1338C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9
1
2
3
4
5
6
7
8
9
outputCopy
1
2
3
4
8
12
5
10
15`
	testutil.AssertEqualCase(t, rawText, 0, CF1338C)
}
