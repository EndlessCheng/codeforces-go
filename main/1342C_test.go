package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1342C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
4 6 5
1 1
1 3
1 5
1 7
1 9
7 10 2
7 8
100 200
outputCopy
0 0 0 2 4 
0 91 `
	testutil.AssertEqualCase(t, rawText, 0, CF1342C)
}
