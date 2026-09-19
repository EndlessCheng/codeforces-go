package main

import (
	"maps"
	"slices"
	"unicode"
)

// https://space.bilibili.com/206214
func braceExpansionII1(expression string) []string {
	type set map[string]struct{}
	i := 0

	var dfs func() set
	dfs = func() set {
		res := set{}
		cur := set{"": {}} // 加个空串，简化后续判断逻辑

		for i < len(expression) {
			ch := expression[i]
			i++

			if ch == '}' { // 归
				break
			}
			if ch == ',' { // 取并集
				maps.Copy(res, cur) // 把 cur 中的字符串都添加到 res 中
				cur = set{"": {}}
			} else if ch == '{' { // 递
				subRes := dfs()
				// 计算 cur 和 subRes 的笛卡尔积
				newSet := make(set, len(cur)*len(subRes)) // 预分配空间
				for s := range cur {
					for t := range subRes {
						newSet[s+t] = struct{}{}
					}
				}
				cur = newSet
			} else { // ch 是字母
				newSet := make(set, len(cur)) // 预分配空间
				t := string(ch)
				for s := range cur {
					newSet[s+t] = struct{}{} // 把 ch 添加到 cur 每个字符串的末尾
				}
				cur = newSet
			}
		}

		maps.Copy(res, cur)
		return res
	}

	return slices.Sorted(maps.Keys(dfs()))
}

func braceExpansionII(expression string) []string {
	type set map[string]struct{}
	type pair struct{ res, cur set }

	stack := []pair{}
	res := set{}
	cur := set{"": {}} // 加个空串，简化后续判断逻辑

	for _, ch := range expression {
		if unicode.IsLower(ch) { // 字母
			newSet := make(set, len(cur)) // 预分配空间
			t := string(ch)
			for s := range cur {
				newSet[s+t] = struct{}{} // 把 ch 添加到 cur 每个字符串的末尾
			}
			cur = newSet
		} else if ch == ',' { // 取并集
			maps.Copy(res, cur) // 把 cur 中的字符串都添加到 res 中
			cur = set{"": {}}
		} else if ch == '{' { // 递
			// 模拟递归
			stack = append(stack, pair{res, cur}) // 递归前，把局部变量 res 和 cur 保存到栈中
			res = set{} // 递归，初始化 res 和 cur
			cur = set{"": {}}
		} else { // 归
			maps.Copy(res, cur)
			subRes := res // 递归结束，返回值为 subRes

			// 从栈中恢复递归之前保存的局部变量
			last := len(stack) - 1
			res = stack[last].res
			cur = stack[last].cur
			stack = stack[:last]

			// 计算 cur 和 subRes 的笛卡尔积
			newSet := make(set, len(cur)*len(subRes)) // 预分配空间
			for s := range cur {
				for t := range subRes {
					newSet[s+t] = struct{}{}
				}
			}
			cur = newSet
		}
	}

	maps.Copy(res, cur)
	return slices.Sorted(maps.Keys(res))
}
