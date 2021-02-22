package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/551/B
// https://codeforces.com/problemset/status/551/problem/B
func TestCF551B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
aaa
a
b
outputCopy
aaa
inputCopy
pozdravstaklenidodiri
niste
dobri
outputCopy
nisteaadddiiklooprrvz
inputCopy
abbbaaccca
ab
aca
outputCopy
ababacabcc`
	testutil.AssertEqualCase(t, rawText, 0, CF551B)
}
