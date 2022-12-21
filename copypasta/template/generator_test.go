package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

// https://codeforces.com/problemset/problem/1293/C
// https://codeforces.com/problemset/status/1291/problem/D
// https://codeforces.com/gym/102253/problem/C
// https://codeforces.com/gym/102253/status/C
func TestGenCodeforcesProblemTemplates(t *testing.T) {
	raw, err := os.ReadFile("data.txt")
	if err != nil {
		t.Fatal(err)
	}
	problemURL := strings.TrimSpace(string(raw))
	if err := GenCodeforcesProblemTemplates(problemURL, true); err != nil {
		t.Fatal(err)
	}
}

// 如果有交互题，赛前先熟悉下交互模板
// 确认下 main.go 和 main_test.go 是否恢复
// https://atcoder.jp/contests/abc262/tasks_print
func TestGenCodeforcesContestTemplates(t *testing.T) {
	const cmdName = CmdAtcoder //
	const contestID = "" //
	const overwrite = false
	rootPath := fmt.Sprintf("../../%s/", contestID)
	for {
		// 配合 https://github.com/xalanq/cf-tool / https://github.com/sempr/cf-tool 使用
		if err := GenCodeforcesContestTemplates(cmdName, rootPath, contestID, overwrite); err != nil {
			//t.Log(err)
		} else {
			break
		}
		time.Sleep(5 * time.Second)
	}
}

// "../../misc/gcj/<year>/<id>/"  需要改模板
// "../../misc/kickstart/<year>/<id>/"  需要改模板
// "../../misc/luogu/contest/<id>/"
// "../../misc/nowcoder/<id>/"
func TestGenTemplates(t *testing.T) {
	const (
		problemNum = 4
		rootPath   = "../../misc/luogu/contest/<id>/"
		overwrite  = false
	)
	if err := GenTemplates(problemNum, rootPath, overwrite); err != nil {
		t.Fatal(err)
	}
}
