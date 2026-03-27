推荐先完成本题的简单版本：[735. 行星碰撞](https://leetcode.cn/problems/asteroid-collision/)，[我的题解](https://leetcode.cn/problems/asteroid-collision/solutions/3938065/yong-zhan-wei-hu-xiao-xing-xing-pythonja-llqq/)。

从左到右遍历这些机器人（需要先按照位置排序），向右的机器人会和向左的机器人碰撞。

遍历到一个向左的机器人时，我们需要找到左边最近的未移除的机器人。这可以用一个栈维护。

如果当前机器人向右，那么直接入栈，继续向后遍历。

如果当前机器人向左，设其健康度为 $h$，栈顶机器人的健康度为 $\textit{top}$，分类讨论：

- 如果 $\textit{top} > h$，那么移除当前机器人，$\textit{top}$ 减一。
- 如果 $\textit{top} = h$，那么两个机器人都移除。
- 如果 $\textit{top} < h$，那么移除栈顶机器人，$h$ 减一。
- 如此循环，直到当前机器人被移除，或者栈顶为空。

⚠**注意**：比大小的这两个健康度都是正整数，所以减一的那个健康度一定大于 $1$。所以减一后，健康度大于 $0$。

代码实现时，直接在 $\textit{healths}$ 上修改，移除机器人 $i$ 相当于把 $\textit{healths}[i]$ 置为 $0$。最后返回 $\textit{healths}$ 中的正数。

```py [sol-Python3]
class Solution:
    def survivedRobotsHealths(self, positions: List[int], healths: List[int], directions: str) -> List[int]:
        # 创建一个下标数组，对下标数组排序，这样不会打乱输入顺序
        idx = sorted(range(len(positions)), key=lambda i: positions[i])

        st = []
        for i in idx:
            if directions[i] == 'R':  # 机器人 i 向右
                st.append(i)
                continue
            while st:  # 栈顶机器人向右
                j = st[-1]
                if healths[j] > healths[i]:  # 栈顶机器人的健康度大
                    healths[i] = 0  # 移除机器人 i
                    healths[j] -= 1
                    break
                if healths[j] == healths[i]:  # 健康度一样大，都移除
                    healths[i] = 0
                    healths[j] = 0
                    st.pop()
                    break
                # 机器人 i 的健康度大
                healths[i] -= 1
                healths[j] = 0  # 移除机器人 j
                st.pop()

        # 返回幸存机器人的健康度
        return [h for h in healths if h > 0]
```

```java [sol-Java]
class Solution {
    public List<Integer> survivedRobotsHealths(int[] positions, int[] healths, String directions) {
        int n = positions.length;
        // 创建一个下标数组，对下标数组排序，这样不会打乱输入顺序
        Integer[] idx = new Integer[n];
        for (int i = 0; i < n; i++) {
            idx[i] = i;
        }
        Arrays.sort(idx, (i, j) -> positions[i] - positions[j]);

        int[] st = new int[n];
        int top = -1;
        for (int i : idx) {
            if (directions.charAt(i) == 'R') { // 机器人 i 向右
                st[++top] = i;
                continue;
            }
            while (top >= 0) { // 栈顶机器人向右
                int j = st[top];
                if (healths[j] > healths[i]) { // 栈顶机器人的健康度大
                    healths[i] = 0; // 移除机器人 i
                    healths[j]--;
                    break;
                }
                if (healths[j] == healths[i]) { // 健康度一样大，都移除
                    healths[i] = 0;
                    healths[j] = 0;
                    top--;
                    break;
                }
                // 机器人 i 的健康度大
                healths[i]--;
                healths[j] = 0; // 移除机器人 j
                top--;
            }
        }

        // 返回幸存机器人的健康度
        List<Integer> ans = new ArrayList<>();
        for (int h : healths) {
            if (h > 0) {
                ans.add(h);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> survivedRobotsHealths(vector<int>& positions, vector<int>& healths, string directions) {
        int n = positions.size();
        // 创建一个下标数组，对下标数组排序，这样不会打乱输入顺序
        vector<int> idx(n);
        ranges::iota(idx, 0); // idx[i] = i
        ranges::sort(idx, {}, [&](int i) { return positions[i]; });

        stack<int> st;
        for (int i : idx) {
            if (directions[i] == 'R') { // 机器人 i 向右
                st.push(i);
                continue;
            }
            while (!st.empty()) { // 栈顶机器人向右
                int j = st.top();
                if (healths[j] > healths[i]) { // 栈顶机器人的健康度大
                    healths[i] = 0; // 移除机器人 i
                    healths[j]--;
                    break;
                }
                if (healths[j] == healths[i]) { // 健康度一样大，都移除
                    healths[i] = 0;
                    healths[j] = 0;
                    st.pop();
                    break;
                }
                // 机器人 i 的健康度大
                healths[i]--;
                healths[j] = 0; // 移除机器人 j
                st.pop();
            }
        }

        // 返回幸存机器人的健康度
        vector<int> ans;
        for (int h : healths) {
            if (h > 0) {
                ans.push_back(h);
            }
        }
        return ans;
    }
};
```

```c [sol-C]
int* _positions;

int cmp(const void* i, const void* j) {
    return _positions[*(int*)i] - _positions[*(int*)j];
}

int* survivedRobotsHealths(int* positions, int positionsSize, int* healths, int healthsSize, char* directions, int* returnSize) {
    int n = positionsSize;
    // 创建一个下标数组，对下标数组排序，这样不会打乱输入顺序
    int* idx = malloc(n * sizeof(int));
    for (int i = 0; i < n; i++) {
        idx[i] = i;
    }
    _positions = positions;
    qsort(idx, n, sizeof(int), cmp);

    int* st = malloc(n * sizeof(int));
    int top = -1;
    for (int k = 0; k < n; k++) {
        int i = idx[k];
        if (directions[i] == 'R') { // 机器人 i 向右
            st[++top] = i;
            continue;
        }
        while (top >= 0) { // 栈顶机器人向右
            int j = st[top];
            if (healths[j] > healths[i]) { // 栈顶机器人的健康度大
                healths[i] = 0; // 移除机器人 i
                healths[j]--;
                break;
            }
            if (healths[j] == healths[i]) { // 健康度一样大，都移除
                healths[i] = 0;
                healths[j] = 0;
                top--;
                break;
            }
            // 机器人 i 的健康度大
            healths[i]--;
            healths[j] = 0; // 移除机器人 j
            top--;
        }
    }

    free(idx);

    // 返回幸存机器人的健康度
    int* ans = st;
    *returnSize = 0;
    for (int i = 0; i < n; i++) {
        if (healths[i] > 0) {
            ans[(*returnSize)++] = healths[i];
        }
    }
    return ans;
}
```

```go [sol-Go]
func survivedRobotsHealths(positions []int, healths []int, directions string) (ans []int) {
	// 创建一个下标数组，对下标数组排序，这样不会打乱输入顺序
	idx := make([]int, len(positions))
	for i := range idx {
		idx[i] = i
	}
	slices.SortFunc(idx, func(i, j int) int { return positions[i] - positions[j] })

	st := []int{}
	for _, i := range idx {
		if directions[i] == 'R' { // 机器人 i 向右
			st = append(st, i)
			continue
		}
		for len(st) > 0 { // 栈顶机器人向右
			j := st[len(st)-1]
			if healths[j] > healths[i] { // 栈顶机器人的健康度大
				healths[i] = 0 // 移除机器人 i
				healths[j]--
				break
			}
			if healths[j] == healths[i] { // 健康度一样大，都移除
				healths[i] = 0
				healths[j] = 0
				st = st[:len(st)-1]
				break
			}
			// 机器人 i 的健康度大
			healths[i]--
			healths[j] = 0 // 移除机器人 j
			st = st[:len(st)-1]
		}
	}

	// 返回幸存机器人的健康度
	for _, h := range healths {
		if h > 0 {
			ans = append(ans, h)
		}
	}
	return
}
```

```js [sol-JavaScript]
var survivedRobotsHealths = function(positions, healths, directions) {
    // 创建一个下标数组，对下标数组排序，这样不会打乱输入顺序
    const idx = Array.from({ length: positions.length }, (_, i) => i)
                     .sort((i, j) => positions[i] - positions[j]);

    const st = [];
    for (const i of idx) {
        if (directions[i] === 'R') { // 机器人 i 向右
            st.push(i);
            continue;
        }
        while (st.length > 0) { // 栈顶机器人向右
            const j = st[st.length - 1];
            if (healths[j] > healths[i]) { // 栈顶机器人的健康度大
                healths[i] = 0; // 移除机器人 i
                healths[j] -= 1;
                break;
            }
            if (healths[j] === healths[i]) { // 健康度一样大，都移除
                healths[i] = 0;
                healths[j] = 0;
                st.pop();
                break;
            }
            // 机器人 i 的健康度大
            healths[i] -= 1;
            healths[j] = 0; // 移除机器人 j
            st.pop();
        }
    }

    // 返回幸存机器人的健康度
    return healths.filter(h => h > 0);
};
```

```rust [sol-Rust]
impl Solution {
    pub fn survived_robots_healths(positions: Vec<i32>, mut healths: Vec<i32>, directions: String) -> Vec<i32> {
        // 创建一个下标数组，对下标数组排序，这样不会打乱输入顺序
        let mut idx = (0..positions.len()).collect::<Vec<_>>();
        idx.sort_unstable_by_key(|&i| positions[i]);

        let directions = directions.as_bytes();
        let mut st = vec![];

        for i in idx {
            if directions[i] == b'R' { // 机器人 i 向右
                st.push(i);
                continue;
            }
            while let Some(&j) = st.last() { // 栈顶机器人向右
                if healths[j] > healths[i] { // 栈顶机器人的健康度大
                    healths[i] = 0; // 移除机器人 i
                    healths[j] -= 1;
                    break;
                }
                if healths[j] == healths[i] { // 健康度一样大，都移除
                    healths[i] = 0;
                    healths[j] = 0;
                    st.pop();
                    break;
                }
                // 机器人 i 的健康度大
                healths[i] -= 1;
                healths[j] = 0; // 移除机器人 j
                st.pop();
            }
        }

        // 返回幸存机器人的健康度
        healths.into_iter().filter(|&h| h > 0).collect()
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{positions}$ 的长度。瓶颈在排序上。虽然我们写了个二重循环，但每个元素至多入栈出栈各一次，所以二重循环的**总**循环次数是 $\mathcal{O}(n)$ 的。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面数据结构题单的「**§3.3 邻项消除**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
