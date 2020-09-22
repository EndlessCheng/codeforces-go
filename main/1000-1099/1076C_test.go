package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1076C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
69
0
1
4
5
999
1000
outputCopy
Y 67.985071301 1.014928699
Y 0.000000000 0.000000000
N
Y 2.000000000 2.000000000
Y 3.618033989 1.381966011
Y 997.998996990 1.001003010
Y 998.998997995 1.001002005`
	testutil.AssertEqualCase(t, rawText, 0, CF1076C)
}
