package __99

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF95A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
bers
ucky
elu
PetrLoveLuckyNumbers
t
outputCopy
PetrLovtTttttNumtttt
inputCopy
4
hello
party
abefglghjdhfgj
IVan
petrsmatchwin
a
outputCopy
petrsmatchwin
inputCopy
2
aCa
cba
abAcaba
c
outputCopy
abCacba
inputCopy
2
A
B
abababBabaBBaBBBBAaaaAAAAA
a
outputCopy
bababaAbabAAbAAAABbbbBBBBB`
	testutil.AssertEqualCase(t, rawText, 0, CF95A)
}
