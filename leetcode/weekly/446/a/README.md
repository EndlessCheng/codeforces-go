按照题意模拟。

如何知道当前位置是否访问过？

创建一个布尔数组 $\textit{vis}$。首次访问 $i$ 时，标记 $\textit{vis}[i]=\texttt{true}$。

如果继续循环，发现 $\textit{vis}[i]$ 等于 $\texttt{true}$，说明我们再次访问了 $i$，退出循环。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1NALczNERr/)，欢迎点赞关注~

## 写法一

```py [sol-Python3]
class Solution:
    def calculateScore(self, instructions: List[str], values: List[int]) -> int:
        n = len(instructions)
        vis = [False] * n
        ans = i = 0
        while 0 <= i < n and not vis[i]:
            vis[i] = True
            if instructions[i][0] == 'a':
                ans += values[i]
                i += 1
            else:
                i += values[i]
        return ans
```

```java [sol-Java]
class Solution {
    public long calculateScore(String[] instructions, int[] values) {
        int n = instructions.length;
        boolean[] vis = new boolean[n];
        long ans = 0;
        int i = 0;
        while (0 <= i && i < n && !vis[i]) {
            vis[i] = true;
            if (instructions[i].charAt(0) == 'a') {
                ans += values[i];
                i++;
            } else {
                i += values[i];
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long calculateScore(vector<string>& instructions, vector<int>& values) {
        int n = instructions.size();
        vector<int> vis(n);
        long long ans = 0;
        int i = 0;
        while (0 <= i && i < n && !vis[i]) {
            vis[i] = true;
            if (instructions[i][0] == 'a') {
                ans += values[i];
                i++;
            } else {
                i += values[i];
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func calculateScore(instructions []string, values []int) (ans int64) {
	n := len(instructions)
	vis := make([]bool, n)
	i := 0
	for 0 <= i && i < n && !vis[i] {
		vis[i] = true
		if instructions[i][0] == 'a' {
			ans += int64(values[i])
			i++
		} else {
			i += values[i]
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{instructions}$ 的长度。每个位置至多访问一次。
- 空间复杂度：$\mathcal{O}(n)$。

## 写法二

用 $\textit{instructions}$ 代替 $\textit{vis}$ 数组：访问到 $i$ 时，把 $\textit{instructions}[i]$ 置为空。

> 注：在实际工程中，不推荐这种会修改入参的写法，除非我们能保证入参不再被使用。

```py [sol-Python3]
class Solution:
    def calculateScore(self, instructions: List[str], values: List[int]) -> int:
        n = len(instructions)
        ans = i = 0
        while 0 <= i < n and instructions[i]:
            s = instructions[i]
            instructions[i] = None
            if s[0] == 'a':
                ans += values[i]
                i += 1
            else:
                i += values[i]
        return ans
```

```java [sol-Java]
class Solution {
    public long calculateScore(String[] instructions, int[] values) {
        int n = instructions.length;
        long ans = 0;
        int i = 0;
        while (0 <= i && i < n && instructions[i] != null) {
            String s = instructions[i];
            instructions[i] = null;
            if (s.charAt(0) == 'a') {
                ans += values[i];
                i++;
            } else {
                i += values[i];
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long calculateScore(vector<string>& instructions, vector<int>& values) {
        int n = instructions.size();
        long long ans = 0;
        int i = 0;
        while (0 <= i && i < n && !instructions[i].empty()) {
            string s = move(instructions[i]);
            if (s[0] == 'a') {
                ans += values[i];
                i++;
            } else {
                i += values[i];
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func calculateScore(instructions []string, values []int) (ans int64) {
	n := len(instructions)
	i := 0
	for 0 <= i && i < n && instructions[i] != "" {
		s := instructions[i]
		instructions[i] = ""
		if s[0] == 'a' {
			ans += int64(values[i])
			i++
		} else {
			i += values[i]
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{instructions}$ 的长度。每个位置至多访问一次。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
