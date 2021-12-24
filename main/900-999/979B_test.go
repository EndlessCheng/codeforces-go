package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/979/B
// https://codeforces.com/problemset/status/979/problem/B
func TestCF979B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
Kuroo
Shiro
Katie
outputCopy
Kuro
inputCopy
7
treasurehunt
threefriends
hiCodeforces
outputCopy
Shiro
inputCopy
1
abcabc
cbabac
ababca
outputCopy
Katie
inputCopy
15
foPaErcvJ
mZaxowpbt
mkuOlaHRE
outputCopy
Draw
inputCopy
60
ddcZYXYbZbcXYcZdYbddaddYaZYZdaZdZZdXaaYdaZZZaXZXXaaZbb
dcdXcYbcaXYaXYcacYabYcbZYdacaYbYdXaccYXZZZdYbbYdcZZZbY
XaZXbbdcXaadcYdYYcbZdcaXaYZabbXZZYbYbcXbaXabcXbXadbZYZ
outputCopy
Draw`
	testutil.AssertEqualCase(t, rawText, -1, CF979B)
}
