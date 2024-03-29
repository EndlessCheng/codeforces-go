// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	t.Log("记得初始化所有全局变量")
	examples := [][3]string{
		{
			`["SORTracker", "add", "add", "get", "add", "get", "add", "get", "add", "get", "add", "get", "get"]`,
			`[[], ["bradford", 2], ["branford", 3], [], ["alps", 2], [], ["orland", 2], [], ["orlando", 3], [], ["alpine", 2], [], []]`,
			`[null, null, null, "branford", null, "alps", null, "bradford", null, "bradford", null, "bradford", "orland"]`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeClassWithExamples(t, Constructor, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-67/problems/sequentially-ordinal-rank-tracker/
