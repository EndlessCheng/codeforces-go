package main

import (
	"sort"
	"strings"
)

/* 字典树+哈希表

我们用一棵字典树去表示整个文件系统，然后 DFS 这棵字典树，并用括号表达式按字典序拼接所有子树。

以示例 $4$ 为例，文件夹 $\texttt{a}$ 的子树的括号表达式为 $\texttt{(x(y))(z)}$。

![](https://assets.leetcode.com/uploads/2021/07/19/lc-dupfolder4_.jpg)

这一表达方式包含了子树的节点名称和父子关系，因此可用于判断两个文件夹是否为相同文件夹。

代码实现时，可以用一个哈希表去存储括号表达式及其对应的文件夹节点列表。在 DFS 结束后，遍历哈希表，若一个括号表达式对应的节点个数大于 $1$，则将该括号表达式对应的所有节点标记为待删除。然后再次 DFS 这棵字典树，仅访问未被删除的节点，并将路径记录到答案中。

*/

// github.com/EndlessCheng/codeforces-go
type folder struct {
	son map[string]*folder
	val string // 文件夹名称
	del bool   // 删除标记
}

func deleteDuplicateFolder(paths [][]string) (ans [][]string) {
	root := &folder{}
	for _, path := range paths {
		// 将 path 加入字典树
		f := root
		for _, s := range path {
			if f.son == nil {
				f.son = map[string]*folder{}
			}
			if f.son[s] == nil {
				f.son[s] = &folder{}
			}
			f = f.son[s]
			f.val = s
		}
	}

	fMap := map[string]*folder{} // 存储括号表达式及其对应的文件夹节点
	var dfs func(*folder) string
	dfs = func(f *folder) string {
		if f.son == nil {
			return "(" + f.val + ")"
		}
		expr := make([]string, 0, len(f.son))
		for _, son := range f.son {
			expr = append(expr, dfs(son))
		}
		sort.Strings(expr)
		subTreeExpr := strings.Join(expr, "") // 按字典序拼接所有子树
		if o := fMap[subTreeExpr]; o != nil {
			o.del = true
			f.del = true
		} else {
			fMap[subTreeExpr] = f
		}
		return "(" + f.val + subTreeExpr + ")"
	}
	dfs(root)

	// 再次 DFS 这棵字典树，仅访问未被删除的节点，并将路径记录到答案中
	path := []string{}
	var dfs2 func(*folder)
	dfs2 = func(f *folder) {
		if f.del {
			return
		}
		path = append(path, f.val)
		ans = append(ans, append([]string(nil), path...))
		for _, son := range f.son {
			dfs2(son)
		}
		path = path[:len(path)-1]
	}
	for _, son := range root.son {
		dfs2(son)
	}
	return
}
