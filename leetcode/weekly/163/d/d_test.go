package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	sampleIns := [][]string{{`[["#","#","#","#","#","#"],["#","T","#","#","#","#"],["#",".",".","B",".","#"],["#",".","#","#",".","#"],["#",".",".",".","S","#"],["#","#","#","#","#","#"]]`}, {`[["#","#","#","#","#","#"],["#","T","#","#","#","#"],["#",".",".","B",".","#"],["#","#","#","#",".","#"],["#",".",".",".","S","#"],["#","#","#","#","#","#"]]`}, {`[["#","#","#","#","#","#"],["#","T",".",".","#","#"],["#",".","#","B",".","#"],["#",".",".",".",".","#"],["#",".",".",".","S","#"],["#","#","#","#","#","#"]]`}, {`[["#","#","#","#","#","#","#"],["#","S","#",".","B","T","#"],["#","#","#","#","#","#","#"]]`}}
	sampleOuts := [][]string{{`3`}, {`-1`}, {`5`}, {`-1`}}
	if err := testutil.RunLeetCodeFunc(t, minPushBox, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/problems/minimum-moves-to-move-a-box-to-their-target-location/
