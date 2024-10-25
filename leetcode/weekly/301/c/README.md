[视频讲解](https://www.bilibili.com/video/BV1aU4y1q7BA)

首先，无论怎么移动，由于 `L` 和 `R` 无法互相穿过对方，那么去掉 `_` 后的剩余字符应该是相同的，否则返回 `false`。

然后用双指针从左向右遍历 $\textit{start}$ 和 $\textit{target}$，分类讨论：

- 如果当前字符为 `L` 且 $i<j$，由于 `L` 由于无法向右移动，返回 `false`；
- 如果当前字符为 `R` 且 $i>j$，由于 `R` 由于无法向左移动，返回 `false`。

遍历完，若中途没有返回 `false` 就返回 `true`。

```py [sol-Python3]
class Solution:
    def canChange(self, start: str, target: str) -> bool:
        if start.replace('_', '') != target.replace('_', ''):
            return False
        j = 0
        for i, c in enumerate(start):
            if c == '_':
                continue
            while target[j] == '_': 
                j += 1
            if i != j and (c == 'L') == (i < j):
                return False
            j += 1
        return True
```

```java [sol-Java]
class Solution {
    public boolean canChange(String start, String target) {
        if (!start.replace("_", "").equals(target.replace("_", ""))) {
            return false;
        }
        for (int i = 0, j = 0; i < start.length(); i++) {
            if (start.charAt(i) == '_') {
                continue;
            }
            while (target.charAt(j) == '_') {
                j++;
            }
            if (i != j && (start.charAt(i) == 'L') == (i < j)) {
                return false;
            }
            j++;
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool canChange(string start, string target) {
        string s = start, t = target;
        s.erase(remove(s.begin(), s.end(), '_'), s.end());
        t.erase(remove(t.begin(), t.end(), '_'), t.end());
        if (s != t) {
            return false;
        }
        for (int i = 0, j = 0; i < start.length(); i++) {
            if (start[i] == '_') {
                continue;
            }
            while (target[j] == '_') {
                j++;
            }
            if (i != j && (start[i] == 'L') == (i < j)) {
                return false;
            }
            j++;
        }
        return true;
    }
};
```

```go [sol-Go]
func canChange(start, target string) bool {
	if strings.ReplaceAll(start, "_", "") != strings.ReplaceAll(target, "_", "") {
		return false
	}
	j := 0
	for i, c := range start {
		if c != '_' {
			for target[j] == '_' {
				j++
			}
			if i != j && c == 'L' == (i < j) {
				return false
			}
			j++
		}
	}
	return true
}
```

```js [sol-JavaScript]
var canChange = function(start, target) {
    if (start.replaceAll('_', '') !== target.replaceAll('_', '')) {
        return false;
    }
    let j = 0;
    for (let i = 0; i < start.length; i++) {
        if (start[i] === '_') {
            continue;
        }
        while (target[j] === '_') {
            j++;
        }
        if (i !== j && (start[i] === 'L') === (i < j)) {
            return false;
        }
        j++;
    }
    return true;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{start}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
