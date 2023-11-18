package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Testb3656(t *testing.T) {
	var cases = [][2]string{
		{
			`10
pop_back 2
push_back 1 1
push_front 1 3
push_front 2 2
push_front 2 3
pop_back 1
size 1
push_back 2 3
back 1
front 1`,
			`1
3
3`,
		},
	}
	testutil.AssertEqualStringCase(t, cases, 0, b3656)
}
