package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	testCases := [][2]string{
		{
			`100
1`,
			`19`,
		},
		{
			`25
2`,
			`14`,
		},
		{
			`314159
2`,
			`937`,
		},
		{
			`9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999
3`,
			`117879300`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
