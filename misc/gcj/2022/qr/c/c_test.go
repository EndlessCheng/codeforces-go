package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"path/filepath"
	"testing"
)

func Test(t *testing.T) {
	dir, _ := filepath.Abs(".")
	t.Logf("Current problem is [%s]", filepath.Base(dir))

	customTestCases := [][2]string{
		{
			`4
4
6 10 12 8
6
5 4 5 4 4 4
10
10 10 7 6 7 4 4 5 7 4
1
10
`,
			`Case #1: 4
Case #2: 5
Case #3: 9
Case #4: 1
`,
		},
	}
	testutil.AssertEqualStringCase(t, customTestCases, 0, run)
}
