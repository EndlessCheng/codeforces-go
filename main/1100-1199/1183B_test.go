package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1183B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
5 1
1 1 2 3 1
4 2
6 4 8 5
2 2
1 6
3 5
5 2 5
outputCopy
2
6
-1
7`
	testutil.AssertEqualCase(t, rawText, 0, CF1183B)
}
