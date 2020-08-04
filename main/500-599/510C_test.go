package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol510C(t *testing.T) {
	// just copy from website
	rawText := `
3
rivest
shamir
adleman
outputCopy
bcdefghijklmnopqrsatuvwxyz
inputCopy
10
tourist
petr
wjmzbmr
yeputons
vepifanov
scottwu
oooooooooooooooo
subscriber
rowdark
tankengineer
outputCopy
Impossible
inputCopy
10
petr
egor
endagorion
feferivan
ilovetanyaromanova
kostka
dmitriyh
maratsnowbear
bredorjaguarturnik
cgyforever
outputCopy
aghjlnopefikdmbcqrstuvwxyz
inputCopy
7
car
care
careful
carefully
becarefuldontforgetsomething
otherwiseyouwillbehacked
goodluck
outputCopy
acbdefhijklmnogpqrstuvwxyz`
	testutil.AssertEqualCase(t, rawText, -1, Sol510C)
}
