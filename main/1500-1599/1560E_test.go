package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1560/problem/E
// https://codeforces.com/problemset/status/1560/problem/E
func TestCF1560E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
abacabaaacaac
nowyouknowthat
polycarppoycarppoyarppyarppyrpprppp
isi
everywherevrywhrvryhrvrhrvhv
haaha
qweqeewew
outputCopy
abacaba bac
-1
polycarp lcoayrp
is si
everywhere ewyrhv
-1
-1`
	testutil.AssertEqualCase(t, rawText, -1, CF1560E)
}
