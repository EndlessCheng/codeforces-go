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

// atc race xxx
// 是否有交互题？熟悉下交互模板
// https://atcoder.jp/contests/abc287/tasks_print
func TestGenCodeforcesContestTemplates(t *testing.T) {
	const cmdName = CmdCodeforces //
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
