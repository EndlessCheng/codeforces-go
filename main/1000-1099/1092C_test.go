package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1092C(t *testing.T) {
	// just copy from website
	rawText := `
5
ba
a
abab
a
aba
baba
ab
aba
outputCopy
SPPSPSPS
inputCopy
3
a
aa
aa
a
outputCopy
PPSS
inputCopy
2
a
c
outputCopy
PS`
	testutil.AssertEqual(t, rawText, CF1092C)
}
