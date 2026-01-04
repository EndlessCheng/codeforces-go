前置题目：[46. 全排列](https://leetcode.cn/problems/permutations/)，视频讲解：[排列型回溯【基础算法精讲 16】](https://www.bilibili.com/video/BV1mY411D7f6/)。

按题目要求模拟即可。

注意为了保证答案有序，可以先对 $\textit{words}$ 从小到大排序，再枚举排列。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def wordSquares(self, words: List[str]) -> List[List[str]]:
        words.sort()  # 保证答案有序
        return [[top, left, right, bottom]
                for top, left, right, bottom in permutations(words, 4)
                if top[0] == left[0] and top[3] == right[0] and
                bottom[0] == left[3] and bottom[3] == right[3]]
```

```java [sol-Java]
class Solution {
    public List<List<String>> wordSquares(String[] words) {
        Arrays.sort(words); // 保证答案有序
        
        int[] path = new int[4];
        boolean[] onPath = new boolean[words.length];
        List<List<String>> ans = new ArrayList<>();

        dfs(words, 0, path, onPath, ans);
        return ans;
    }

    private void dfs(String[] words, int i, int[] path, boolean[] onPath, List<List<String>> ans) {
        if (i == 4) {
            String top = words[path[0]];
            String left = words[path[1]];
            String right = words[path[2]];
            String bottom = words[path[3]];
            if (top.charAt(0) == left.charAt(0) && top.charAt(3) == right.charAt(0)
                    && bottom.charAt(0) == left.charAt(3) && bottom.charAt(3) == right.charAt(3)) {
                ans.add(Arrays.asList(top, left, right, bottom));
            }
            return;
        }

        for (int j = 0; j < words.length; j++) {
            if (!onPath[j]) {
                path[i] = j; // 从没有选的下标中选一个
                onPath[j] = true; // 已选上
                dfs(words, i + 1, path, onPath, ans);
                onPath[j] = false; // 恢复现场
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<string>> wordSquares(vector<string>& words) {
        ranges::sort(words); // 保证答案有序

        int path[4];
        vector<int8_t> on_path(words.size());
        vector<vector<string>> ans;

        auto dfs = [&](this auto&& dfs, int i) -> void {
            if (i == 4) {
                string& top = words[path[0]];
                string& left = words[path[1]];
                string& right = words[path[2]];
                string& bottom = words[path[3]];
                if (top[0] == left[0] && top[3] == right[0] && bottom[0] == left[3] && bottom[3] == right[3]) {
                    ans.push_back({top, left, right, bottom});
                }
                return;
            }

            for (int j = 0; j < words.size(); j++) {
                if (!on_path[j]) {
                    path[i] = j; // 从没有选的下标中选一个
                    on_path[j] = true; // 已选上
                    dfs(i + 1);
                    on_path[j] = false; // 恢复现场
                }
            }
        };

        dfs(0);
        return ans;
    }
};
```

```go [sol-Go]
func wordSquares(words []string) (ans [][]string) {
	slices.Sort(words) // 保证答案有序

	path := [4]int{}
	onPath := make([]bool, len(words))

	var dfs func(int)
	dfs = func(i int) {
		if i == 4 {
			top := words[path[0]]
			left := words[path[1]]
			right := words[path[2]]
			bottom := words[path[3]]
			if top[0] == left[0] && top[3] == right[0] && bottom[0] == left[3] && bottom[3] == right[3] {
				ans = append(ans, []string{top, left, right, bottom})
			}
			return
		}

		for j, on := range onPath {
			if !on {
				path[i] = j      // 从没有选的下标中选一个
				onPath[j] = true // 已选上
				dfs(i + 1)
				onPath[j] = false // 恢复现场
			}
		}
	}

	dfs(0)
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(P_n^4) = \mathcal{O}(n^4)$，其中 $n$ 是 $\textit{words}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。排序的栈开销和返回值不计入。

## 专题训练

见下面回溯题单的「**§4.5 排列型回溯**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
