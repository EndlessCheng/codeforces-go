package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1367D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
abac
3
2 1 0
abc
1
0
abba
3
1 0 1
ecoosdcefr
10
38 13 24 14 11 5 3 24 17 0
outputCopy
aac
b
aba
codeforces`
	testutil.AssertEqualCase(t, rawText, 0, CF1367D)
}
