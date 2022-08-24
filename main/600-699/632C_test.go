package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol623C(t *testing.T) {
	// just copy from website
	rawText := `
2
bbz
b
outputCopy
bbbz
inputCopy
2
b
bbz
outputCopy
bbbz
inputCopy
2
b
bba
outputCopy
bbab
inputCopy
2
b
bza
outputCopy
bbza
inputCopy
2
b
baa
outputCopy
baab
inputCopy
2
a
az
outputCopy
aaz
inputCopy
4
abba
abacaba
bcd
er
outputCopy
abacabaabbabcder
inputCopy
5
x
xx
xxa
xxaa
xxaaa
outputCopy
xxaaaxxaaxxaxxx
inputCopy
3
c
cb
cba
outputCopy
cbacbc`
	testutil.AssertEqualCase(t, rawText, 0, Sol632C)
}
