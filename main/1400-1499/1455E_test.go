package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1455E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
0 2
4 2
2 0
2 4
1 0
2 0
4 0
6 0
1 6
2 2
2 5
4 1
outputCopy
8
7
5`
	testutil.AssertEqualCase(t, rawText, 0, CF1455E)
}
