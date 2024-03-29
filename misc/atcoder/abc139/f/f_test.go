// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [f]")
	testCases := [][2]string{
		{
			`3
0 10
5 -5
-5 -5`,
			`10.000000000000000000000000000000000000000000000000`,
		},
		{
			`5
1 1
1 0
0 1
-1 0
0 -1`,
			`2.828427124746190097603377448419396157139343750753`,
		},
		{
			`5
1 1
2 2
3 3
4 4
5 5`,
			`21.213203435596425732025330863145471178545078130654`,
		},
		{
			`3
0 0
0 1
1 0`,
			`1.414213562373095048801688724209698078569671875376`,
		},
		{
			`1
90447 91000`,
			`128303.000000000000000000000000000000000000000000000000`,
		},
		{
			`2
96000 -72000
-72000 54000`,
			`120000.000000000000000000000000000000000000000000000000`,
		},
		{
			`10
1 2
3 4
5 6
7 8
9 10
11 12
13 14
15 16
17 18
19 20`,
			`148.660687473185055226120082139313966514489855137208`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
	//testutil.AssertEqualRunResults(t, testCases, 0, runAC, run)
}
// https://atcoder.jp/contests/abc139/tasks/abc139_f
