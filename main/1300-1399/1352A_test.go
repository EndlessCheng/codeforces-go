package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1352A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5009
7
9876
10000
10
outputCopy
2
5000 9
1
7 
4
800 70 6 9000 
1
10000 
1
10`
	testutil.AssertEqualCase(t, rawText, 0, CF1352A)
}
