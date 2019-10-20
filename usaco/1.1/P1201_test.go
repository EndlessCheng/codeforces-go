package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSolP1201(t *testing.T) {
	inputs := []string{
		`5
dave
laura
owen
vick
amr
dave
200 3
laura
owen
vick
owen
500 1
dave
amr
150 2
vick
owen
laura
0 2
amr
vick
vick
0 0`,
	}
	answers := []string{
		`dave 302
laura 66
owen -359
vick 141
amr -150`,
	}
	testutil.AssertEqualStringCase(t, inputs, answers, 0, SolP1201)
}
