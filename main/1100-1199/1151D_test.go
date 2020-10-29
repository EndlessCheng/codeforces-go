package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1151D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4 2
2 3
6 1
outputCopy
12
inputCopy
4
2 4
3 3
7 1
2 3
outputCopy
25
inputCopy
10
5 10
12 4
31 45
20 55
30 17
29 30
41 32
7 1
5 5
3 15
outputCopy
1423`
	testutil.AssertEqualCase(t, rawText, 0, CF1151D)
}
