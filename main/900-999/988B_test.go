package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF988B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
a
aba
abacaba
ba
aba
outputCopy
YES
a
ba
aba
aba
abacaba
inputCopy
5
a
abacaba
ba
aba
abab
outputCopy
NO
inputCopy
3
qwerty
qwerty
qwerty
outputCopy
YES
qwerty
qwerty
qwerty`
	testutil.AssertEqualCase(t, rawText, 0, CF988B)
}
