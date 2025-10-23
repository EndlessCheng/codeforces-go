package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://loj.ac/p/162
func Test_p162(t *testing.T) {
	testCases := [][2]string{
		{
			`2 3
1 2 3`,
			`2 4 8`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, p162)
}
