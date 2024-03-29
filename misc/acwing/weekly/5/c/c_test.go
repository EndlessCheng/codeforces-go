// Code generated by copypasta/template/acwing/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{
			`3
2 3
1 1
3 2
3 2 3
3 2 3`,
			`8
3
1 2 3 
0`,
		},
		{
			`3
2 1
1 2
3 3
23 2 23
3 2 3`,
			`27
1
2 
2
1 2
2 3`,
		},
		{
			`3
1 1
1 1
1 1
3 2 3
3 2 3`,
			`2
1
2 
2
1 2
1 3`,
		},
	}
	target := 0 // -1
	testutil.AssertEqualStringCase(t, testCases, target, run)
}
