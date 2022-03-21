package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1654/B
// https://codeforces.com/problemset/status/1654/problem/B
func TestCF1654B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
abcabdc
a
bbbbbbbbbb
codeforces
cffcfccffccfcffcfccfcffccffcfccf
zyzyzwxxyyxxyyzzyzzxxwzxwywxwzxxyzzw
outputCopy
abdc
a
b
deforces
cf
xyzzw`
	testutil.AssertEqualCase(t, rawText, 0, CF1654B)
}
