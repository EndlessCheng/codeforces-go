package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF827A(t *testing.T) {
	// just copy from website
	rawText := `
3
a 4 1 3 5 7
ab 2 1 5
ca 1 4
outputCopy
abacaba
inputCopy
1
a 1 3
outputCopy
aaa
inputCopy
3
ab 1 1
aba 1 3
ab 2 3 5
outputCopy
ababab`
	testutil.AssertEqualCase(t, rawText, 0, CF827A)
}
