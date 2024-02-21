package main

import (
	"fmt"
	"testing"
	"time"
)

// atc race abc323
// 是否有交互题？熟悉下交互模板
// https://atcoder.jp/contests/abc323/tasks_print
// https://atcoder.jp/contests/abc323/submit
func TestGenCodeforcesContestTemplates(t *testing.T) {
	const cmdName = CmdCodeforces //
	const contestID = "" //
	const overwrite = false
	rootPath := fmt.Sprintf("../../%s/", contestID)
	for {
		// 配合 https://github.com/xalanq/cf-tool 或 https://github.com/sempr/cf-tool 使用
		if err := GenCodeforcesContestTemplates(cmdName, rootPath, contestID, overwrite); err != nil {
			//t.Log(err)
		} else {
			break
		}
		time.Sleep(5 * time.Second)
	}
}
