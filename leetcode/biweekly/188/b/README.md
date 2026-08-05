栅栏的高度有两种情况：

1. 原始木板的高度。
2. 两块原始木板合成后的高度。

如果枚举高度 $h$，再统计最多有多少块木板可以组成 $h$（例如用相向双指针），时间复杂度是 $\mathcal{O}(n^3)$ 的，太慢了。

反过来，枚举所有木板对 $(\textit{planks}[i], \textit{planks}[j])$，用哈希表记录高度 $h = \textit{planks}[i] + \textit{planks}[j]$ 新增了多少块木板，就可以在枚举高度 $h$ 的时候，$\mathcal{O}(1)$ 知道最多有多少块木板可以组成 $h$。

为此，需要用两个哈希表：

1. 用哈希表 $\textit{cnt}$ 统计 $\textit{planks}$ 每个元素（高度）的出现次数。
2. 用哈希表 $\textit{cntPair}$ 统计合成木板的高度的出现次数。枚举 $\textit{cnt}$ 中的木板对 $(x,y)$，其中 $x<y$，这两种木板最多可以组成 $\min(\textit{cnt}[x], \textit{cnt}[y])$ 块高度为 $x+y$ 的合成木板。此外，高度为 $x$ 的木板可以内部两两组合，最多可以组成 $\left\lfloor\dfrac{\textit{cnt}[x]}{2}\right\rfloor$ 块高度为 $2x$ 的合成木板。

> **注**：对于一个固定的 $h$，木板 $x$ 只能与唯一的互补值 $h-x$ 配对，所以不同的木板对之间不会重复使用同一种木板。

最后，遍历 $\textit{cnt}$ 和 $\textit{cntPair}$ 中的高度 $h$，那么最多有

$$
\textit{cnt}[h] + \textit{cntPair}[h] 
$$

块高为 $h$ 的木板。

[本题视频讲解](https://www.bilibili.com/video/BV1Y73R6zEwG/?t=1m10s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maximumWidth(self, planks: list[int]) -> int:
        # 统计 planks 的元素出现次数
        cnt = Counter(planks)

        # 枚举所有高度对 (x,y)
        ans = 0
        cnt_pair = defaultdict(int)
        for x, c in cnt.items():
            cnt_pair[x] += c  # 方便最后算 max
            cnt_pair[x * 2] += c // 2  # 高为 x 的木板内部配对
            for y, c2 in cnt.items():
                if y > x:  # 避免 x+y 和 y+x 重复统计
                    cnt_pair[x + y] += min(c, c2)

        # 枚举栅栏高度
        return max(cnt_pair.values())
```

```java [sol-Java]
class Solution {
    public int maximumWidth(int[] planks) {
        // 统计 planks 的元素出现次数
        HashMap<Integer, Integer> cnt = new HashMap<>();
        for (int x : planks) {
            cnt.merge(x, 1, Integer::sum); // cnt[x]++
        }

        // 枚举所有高度对 (x,y)
        HashMap<Integer, Integer> cntPair = new HashMap<>();
        for (Map.Entry<Integer, Integer> e : cnt.entrySet()) {
            int x = e.getKey();
            int c = e.getValue();
            cntPair.merge(x, c, Integer::sum); // 方便最后算 max
            cntPair.merge(x * 2, c / 2, Integer::sum); // 高为 x 的木板内部配对
            for (Map.Entry<Integer, Integer> e2 : cnt.entrySet()) {
                int y = e2.getKey();
                int c2 = e2.getValue();
                if (y > x) { // 避免 x+y 和 y+x 重复统计
                    cntPair.merge(x + y, Math.min(c, c2), Integer::sum);
                }
            }
        }

        // 枚举栅栏高度
        return Collections.max(cntPair.values());
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumWidth(vector<int>& planks) {
        // 统计 planks 的元素出现次数
        unordered_map<int, int> cnt;
        for (int x : planks) {
            cnt[x]++;
        }

        // 枚举所有高度对 (x,y)
        unordered_map<int, int> cnt_pair;
        for (auto& [x, c] : cnt) {
            cnt_pair[x] += c; // 方便最后算 max
            cnt_pair[x * 2] += c / 2; // 高为 x 的木板内部配对
            for (auto& [y, c2] : cnt) {
                if (y > x) { // 避免 x+y 和 y+x 重复统计
                    cnt_pair[x + y] += min(c, c2);
                }
            }
        }

        // 枚举栅栏高度
        int ans = 0;
        for (auto& [_, c] : cnt_pair) {
            ans = max(ans, c);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumWidth(planks []int) (ans int) {
	// 统计 planks 的元素出现次数
	cnt := map[int]int{}
	for _, x := range planks {
		cnt[x]++
	}

	// 枚举所有高度对 (x,y)
	cntPair := map[int]int{}
	for x, c := range cnt {
		cntPair[x] += c // 方便最后算 max
		cntPair[x*2] += c / 2 // 高为 x 的木板内部配对
		for y, c2 := range cnt {
			if y > x { // 避免 x+y 和 y+x 重复统计
				cntPair[x+y] += min(c, c2)
			}
		}
	}

	// 枚举栅栏高度
	for _, c := range cntPair {
		ans = max(ans, c)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{planks}$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2)$。

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
