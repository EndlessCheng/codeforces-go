package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF598D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 6 3
******
*..*.*
******
*....*
******
2 2
2 5
4 3
outputCopy
6
4
10
inputCopy
4 4 1
****
*..*
*.**
****
3 2
outputCopy
8`
	testutil.AssertEqualCase(t, rawText, 0, CF598D)
}
