package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF120H(t *testing.T) {
	// just copy from website
	rawText := `
6
privet
spasibo
codeforces
java
marmelad
normalno
outputCopy
e
ab
c
av
aa
a
inputCopy
5
aaa
aa
a
aaaa
aaaaa
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF120H)
}
