// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1907/problem/B
// https://codeforces.com/problemset/status/1907/problem/B
func Test_cf1907B(t *testing.T) {
	testCases := [][2]string{
		{
			`12
ARaBbbitBaby
YetAnotherBrokenKeyboard
Bubble
Improbable
abbreviable
BbBB
BusyasaBeeinaBedofBloomingBlossoms
CoDEBARbIES
codeforces
bobebobbes
b
TheBBlackbboard`,
			`ity
YetnotherrokenKeoard
le
Imprle
revile

usyasaeeinaedofloominglossoms
CDARIES
codeforces
es

helaoard`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1907B)
}
