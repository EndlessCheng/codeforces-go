package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1045I(t *testing.T) {
	// just copy from website
	rawText := `
3
aa
bb
cd
outputCopy
1
inputCopy
6
aab
abcac
dffe
ed
aa
aade
outputCopy
6
inputCopy
20
iw
ix
udb
bg
oi
uo
jsm
um
s
quy
qo
bxct
ng
rmr
nu
ps
io
kh
w
k
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, 0, CF1045I)
}
