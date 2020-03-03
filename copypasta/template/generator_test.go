package main

import (
	"testing"
)

// 生成比赛模板（需要先在 dash 中创建对应目录）
func TestGenContestTemplates(t *testing.T) {
	// TODO 勿扰 & 添加闹钟：剩余 10 分钟和剩余 5 分钟时提醒
	const contestID = "1321"
	const overwrite = false
	if err := GenContestTemplates(contestID, overwrite); err != nil {
		t.Fatal(err)
	}
}

// 生成单道题目的模板（Codeforces）
// https://codeforces.com/problemset/problem/1293/C
// https://codeforces.com/problemset/status/1291/problem/D
// https://codeforces.com/gym/102253/problem/C
func TestGenCodeforcesNormalTemplates(t *testing.T) {
	const problemURL = "https://codeforces.ml/problemset/problem/1301/D"
	if err := GenCodeforcesNormalTemplates(problemURL); err != nil {
		t.Fatal(err)
	}
}

// 生成单道题目的模板（非 Codeforces）
func TestGenNormalTemplates(t *testing.T) {
	const rootPath = "../../nowcoder/2720/"
	const overwrite = false
	if err := GenNormalTemplates(rootPath, overwrite); err != nil {
		t.Fatal(err)
	}
}
