先 DFS 算出树的高度 $h$，做法同 [104. 二叉树的最大深度](https://leetcode.cn/problems/maximum-depth-of-binary-tree/)。

再写一个自顶向下的 DFS，从 $h$ 开始，每往下走一步就把 $h$ 减一，就是题目中的 $h-d+1$ 了。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def weightedSum(self, parent: list[int], nums: list[int]) -> int:
        n = len(parent)
        g = [[] for _ in range(n)]
        for i in range(1, n):
            g[parent[i]].append(i)

        def get_h(x: int, fa: int) -> int:
            h = 0
            for y in g[x]:
                if y != fa:
                    h = max(h, get_h(y, x))
            return h + 1
        h = get_h(0, -1)

        def dfs(x: int, fa: int, weight: int):
            nonlocal ans
            ans += nums[x] * weight
            for y in g[x]:
                if y != fa:
                    dfs(y, x, weight - 1)
        ans = 0
        dfs(0, -1, h)
        return ans
```

```java [sol-Java]
class Solution {
    public long weightedSum(int[] parent, int[] nums) {
        int n = parent.length;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        for (int i = 1; i < n; i++) {
            g[parent[i]].add(i);
        }

        int h = getH(0, -1, g);
        return dfs(0, -1, h, g, nums);
    }

    private int getH(int x, int fa, List<Integer>[] g) {
        int h = 0;
        for (int y : g[x]) {
            if (y != fa) {
                h = Math.max(h, getH(y, x, g));
            }
        }
        return h + 1;
    }

    private long dfs(int x, int fa, int weight, List<Integer>[] g, int[] nums) {
        long ans = (long) nums[x] * weight;
        for (int y : g[x]) {
            if (y != fa) {
                ans += dfs(y, x, weight - 1, g, nums);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long weightedSum(vector<int>& parent, vector<int>& nums) {
        int n = parent.size();
        vector<vector<int>> g(n);
        for (int i = 1; i < n; i++) {
            g[parent[i]].push_back(i);
        }

        auto get_h = [&](this auto&& get_h, int x, int fa) -> int {
            int h = 0;
            for (int y : g[x]) {
                if (y != fa) {
                    h = max(h, get_h(y, x));
                }
            }
            return h + 1;
        };
        int h = get_h(0, -1);

        long long ans = 0;
        auto dfs = [&](this auto&& dfs, int x, int fa, int weight) -> void {
            ans += 1LL * nums[x] * weight;
            for (int y : g[x]) {
                if (y != fa) {
                    dfs(y, x, weight - 1);
                }
            }
        };
        dfs(0, -1, h);
        return ans;
    }
};
```

```go [sol-Go]
func weightedSum(parent []int, nums []int) (ans int64) {
	n := len(parent)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parent[i]
		g[p] = append(g[p], i)
	}

	var getH func(int, int) int
	getH = func(x, fa int) (h int) {
		for _, y := range g[x] {
			if y != fa {
				h = max(h, getH(y, x))
			}
		}
		return h + 1
	}
	h := getH(0, -1)

	var dfs func(int, int, int)
	dfs = func(x, fa, weight int) {
		ans += int64(nums[x]) * int64(weight)
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x, weight-1)
			}
		}
	}
	dfs(0, -1, h)
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面树题单的「**§3.2 自顶向下 DFS**」和「**§3.3 自底向上 DFS**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
