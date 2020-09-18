package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1311C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4 2
abca
1 3
10 5
codeforces
2 8 3 2 9
26 10
qwertyuioplkjhgfdsazxcvbnm
20 10 1 2 3 5 10 5 9 4
outputCopy
4 2 2 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 
0 0 9 4 5 3 0 0 0 0 0 0 0 0 9 0 0 3 1 0 0 0 0 0 0 0 
2 1 1 2 9 2 2 2 5 2 2 2 1 1 5 4 11 8 2 7 5 1 10 1 5 2 `
	testutil.AssertEqualCase(t, rawText, 0, CF1311C)
}
