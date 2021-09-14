package main

import "sort"

/* 简洁写法

将角色按照攻击**从大到小**排序，攻击相同的按照防御**从小到大**排序。

然后遍历数组，维护遍历过的角色的防御的最大值 $\textit{maxDef}$。由于攻击已经按照从大到小排序了，对于当前角色 $p$，如果 $p$ 的防御小于 $\textit{maxDef}$，那么说明前面有防御比 $p$ 高的角色（记作 $q$）；同时，对于攻击相同的角色，由于我们是按照防御从小到大排序的，所以不会出现 $q$ 的防御比 $p$ 高，但是 $q$ 的攻击和 $p$ 相同的情况，因此 $q$ 的攻击必然大于 $p$，这样 $p$ 就是一个弱角色。

*/

// github.com/EndlessCheng/codeforces-go
func numberOfWeakCharacters(a [][]int) (ans int) {
	sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a[0] > b[0] || a[0] == b[0] && a[1] < b[1] })
	maxDef := 0
	for _, p := range a {
		if p[1] < maxDef {
			ans++
		} else {
			maxDef = p[1]
		}
	}
	return
}
