package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF724D(t *testing.T) {
	// just copy from website
	rawText := `
3
cbabc
outputCopy
a
inputCopy
2
abcab
outputCopy
aab
inputCopy
3
bcabcbaccba
outputCopy
aaabb
inputCopy
5
immaydobun
outputCopy
ab
inputCopy
5
wjjdqawypvtgrncmqvcsergermprauyevcegjtcrrblkwiugrcjfpjyxngyryxntauxlouvwgjzpsuxyxvhavgezwtuzknetdibv
outputCopy
aaaabbcccccddeeeeeefggggggghiijjjjjjkkllmmnnnnoppppqqrrrrrrrrsstttttu
`
	testutil.AssertEqualCase(t, rawText, 0, CF724D)
}
