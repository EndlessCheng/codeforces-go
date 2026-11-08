如果你从未做过类似题目，推荐：

1. 先做一道基础的邻项消除问题：[3561. 移除相邻字符](https://leetcode.cn/problems/resulting-string-after-adjacent-removals/)。
2. 再做一道和本题非常像的题目：[1209. 删除字符串中的所有相邻重复项 II](https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string-ii/)。

仿照 1209 的思路，栈中维护三元组 $(类型,颜色,出现次数)$。

设栈顶的类型为 $t$，出现次数为 $\textit{cnt}$。如果恰好满足 $\textit{groupSize}[t] = \textit{cnt}$，那么移除栈顶。

遍历结束后，如果栈为空，说明我们移除了所有元素。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def canRemoveAll(self, types: list[int], colors: list[int], groupSize: list[int]) -> bool:
        st = []  # [类型, 颜色, 出现次数]

        for type, color in zip(types, colors):
            if st and type == st[-1][0] and color == st[-1][1]:
                st[-1][2] += 1
            else:
                st.append([type, color, 1])
            # 只要发现满足要求的连续元素，就立刻移除
            if groupSize[st[-1][0]] == st[-1][2]:
                st.pop()

        return not st
```

```java [sol-Java]
class Solution {
    public boolean canRemoveAll(int[] types, int[] colors, int[] groupSize) {
        Deque<int[]> st = new ArrayDeque<>();

        for (int i = 0; i < types.length; i++) {
            int tp = types[i];
            int color = colors[i];
            if (!st.isEmpty() && tp == st.peek()[0] && color == st.peek()[1]) {
                st.peek()[2]++;
            } else {
                st.push(new int[]{tp, color, 1});
            }
            // 只要发现满足要求的连续元素，就立刻移除
            if (groupSize[st.peek()[0]] == st.peek()[2]) {
                st.pop();
            }
        }

        return st.isEmpty();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool canRemoveAll(vector<int>& types, vector<int>& colors, vector<int>& groupSize) {
        stack<array<int, 3>> st;

        for (int i = 0; i < types.size(); i++) {
            int tp = types[i], color = colors[i];
            if (!st.empty() && tp == st.top()[0] && color == st.top()[1]) {
                st.top()[2]++;
            } else {
                st.push({tp, color, 1});
            }
            // 只要发现满足要求的连续元素，就立刻移除
            if (groupSize[st.top()[0]] == st.top()[2]) {
                st.pop();
            }
        }

        return st.empty();
    }
};
```

```go [sol-Go]
func canRemoveAll(types, colors, groupSize []int) bool {
	type tuple struct{ tp, color, cnt int }
	st := []tuple{}
	for i, tp := range types {
		color := colors[i]
		if len(st) > 0 && tp == st[len(st)-1].tp && color == st[len(st)-1].color {
			st[len(st)-1].cnt++
		} else {
			st = append(st, tuple{tp, color, 1})
		}
		// 只要发现满足要求的连续元素，就立刻移除
		if groupSize[st[len(st)-1].tp] == st[len(st)-1].cnt {
			st = st[:len(st)-1]
		}
	}
	return len(st) == 0
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{types}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面数据结构题单的「**§3.3 邻项消除**」。

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
