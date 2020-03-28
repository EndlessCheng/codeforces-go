package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1183G(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
8
1 0
4 1
2 0
4 1
5 1
6 1
3 0
2 0
4
1 1
1 1
2 1
2 1
9
2 0
2 0
4 1
4 1
4 1
7 0
7 1
7 0
7 1
outputCopy
3 3
3 3
9 5`
	testutil.AssertEqualCase(t, rawText, 0, CF1183G)
}
