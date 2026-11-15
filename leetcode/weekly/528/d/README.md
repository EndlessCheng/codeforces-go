**前置知识**：

1. [LogTrick 入门教程](https://zhuanlan.zhihu.com/p/1933215367158830792)。
2. [单调栈【基础算法精讲 26】](https://www.bilibili.com/video/BV1VN411J7S7/)。

本题需要判断子数组的与值 $\textit{andVal}$ 是否为子数组的最小值。

我们可以在遍历 $\textit{nums}$ 的同时，维护一个单调递增的栈 $\textit{st}$，保存 $\textit{nums}$ 的下标。相同的 $\textit{nums}[i]$ 保留最右边的。

在单调栈中二分查找 $\textit{andVal}$，假设结果为 $\textit{st}[j]$，那么子数组左端点最大为 $\textit{st}[j]$，最小为 $\textit{st}[j-1]+1$。把这个范围与 LogTrick 的范围求交集，即为子数组左端点的范围。把范围的大小加到答案中。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

## 优化前

```py [sol-Python3]
class Solution:
    def countConsistent(self, nums: list[int]) -> int:
        data = []
        st = []
        ans = 0

        for i, x in enumerate(nums):
            # x 入单调栈
            while st and nums[st[-1]] >= x:
                st.pop()
            st.append(i)

            # 计算右端点为 i 的子数组与值
            for p in data:
                p[0] &= x
            # x 单独一个数作为子数组
            data.append([x, i])

            # 原地去重（相同与值只保留最左边的）
            idx = 1
            for j in range(1, len(data)):
                if data[j][0] != data[j - 1][0]:
                    data[idx] = data[j]
                    idx += 1
            del data[idx:]

            for k, (and_val, left) in enumerate(data):
                # 判断子数组最小值是否为 andVal
                j = bisect_left(st, and_val, key=lambda j: nums[j])
                if j == len(st) or nums[st[j]] != and_val:
                    continue

                right = data[k + 1][1] - 1 if k < len(data) - 1 else i
                right = min(right, st[j])
                # 注：由于 nums[st[j-1]] < and_val，假如子数组包含 st[j-1]，那么与值会比 and_val 还小，矛盾
                # 所以子数组不可能包含 st[j-1]，所以无需考虑 st[j-1]
                ans += max(right - left + 1, 0)

        return ans
```

```java [sol-Java]
class Solution {
    public long countConsistent(int[] nums) {
        List<int[]> data = new ArrayList<>();
        List<Integer> st = new ArrayList<>();
        long ans = 0;

        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];

            // x 入单调栈
            while (!st.isEmpty() && nums[st.getLast()] >= x) {
                st.removeLast();
            }
            st.add(i);

            // 计算右端点为 i 的子数组与值
            for (int[] p : data) {
                p[0] &= x;
            }
            // x 单独一个数作为子数组
            data.add(new int[]{x, i});

            // 原地去重（相同与值只保留最左边的）
            int m = 1;
            for (int j = 1; j < data.size(); j++) {
                if (data.get(j)[0] != data.get(j - 1)[0]) {
                    data.set(m++, data.get(j));
                }
            }
            data.subList(m, data.size()).clear();

            for (int k = 0; k < m; k++) {
                int andVal = data.get(k)[0];
                // 判断子数组最小值是否为 andVal
                int j = lowerBound(st, nums, andVal);
                if (j == st.size() || nums[st.get(j)] != andVal) {
                    continue;
                }

                int left = data.get(k)[1];
                int right = k < m - 1 ? data.get(k + 1)[1] - 1 : i;
                right = Math.min(right, st.get(j));
                // 注：由于 nums[st[j-1]] < andVal，假如子数组包含 st[j-1]，那么与值会比 andVal 还小，矛盾
                // 所以子数组不可能包含 st[j-1]，所以无需考虑 st[j-1]
                ans += Math.max(right - left + 1, 0);
            }
        }

        return ans;
    }

    private int lowerBound(List<Integer> st, int[] nums, int target) {
        int low = -1;
        int high = st.size();
        while (low + 1 < high) {
            int mid = (low + high) >>> 1;
            if (nums[st.get(mid)] >= target) {
                high = mid;
            } else {
                low = mid;
            }
        }
        return high;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countConsistent(vector<int>& nums) {
        vector<pair<int, int>> data;
        vector<int> st;
        long long ans = 0;

        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];

            // x 入单调栈
            while (!st.empty() && nums[st.back()] >= x) {
                st.pop_back();
            }
            st.push_back(i);

            // 计算右端点为 i 的子数组与值
            for (auto& [and_val, _] : data) {
                and_val &= x;
            }
            // x 单独一个数作为子数组
            data.emplace_back(x, i);

            // 原地去重（相同与值只保留最左边的）
            int m = 1;
            for (int j = 1; j < data.size(); j++) {
                if (data[j].first != data[j - 1].first) {
                    data[m++] = data[j];
                }
            }
            data.resize(m);

            for (int k = 0; k < m; k++) {
                auto& [and_val, left] = data[k];
                // 判断子数组最小值是否为 andVal
                int j = ranges::lower_bound(st, and_val, {}, [&](int j) { return nums[j]; }) - st.begin();
                if (j == st.size() || nums[st[j]] != and_val) {
                    continue;
                }

                int right = k + 1 < m ? data[k + 1].second - 1 : i;
                right = min(right, st[j]);
                // 注：由于 nums[st[j-1]] < and_val，假如子数组包含 st[j-1]，那么与值会比 and_val 还小，矛盾
                // 所以子数组不可能包含 st[j-1]，所以无需考虑 st[j-1]
                ans += max(right - left + 1, 0);
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
func countConsistent(nums []int) (ans int64) {
	type pair struct{ and, left int } // 子数组与值，最小左端点
	data := []pair{}
	st := []int{}

	for i, x := range nums {
		// x 入单调栈
		for len(st) > 0 && nums[st[len(st)-1]] >= x {
			st = st[:len(st)-1]
		}
		st = append(st, i)

		// 计算右端点为 i 的子数组与值
		for j := range data {
			data[j].and &= x
		}
		// x 单独一个数作为子数组
		data = append(data, pair{x, i})

		// 原地去重（相同与值只保留最左边的）
		idx := 1
		for j := 1; j < len(data); j++ {
			if data[j].and != data[j-1].and {
				data[idx] = data[j]
				idx++
			}
		}
		data = data[:idx]

		m := len(st)
		for k, p := range data {
			andVal := p.and
			// 判断子数组最小值是否为 andVal
			j := sort.Search(m, func(j int) bool { return nums[st[j]] >= andVal })
			if j == m || nums[st[j]] != andVal {
				continue
			}

			right := i
			if k < len(data)-1 {
				right = data[k+1].left - 1
			}
			right = min(right, st[j])
			// 注：由于 nums[st[j-1]] < andVal，假如子数组包含 st[j-1]，那么与值会比 andVal 还小，矛盾
			// 所以子数组不可能包含 st[j-1]，所以无需考虑 st[j-1]
			ans += int64(max(right-p.left+1, 0))
		}
	}

	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(n)$。

## 优化

额外用一个哈希表记录 $\textit{nums}[i]$ 的下标，这样可以 $\mathcal{O}(1)$ 获取 $\textit{andVal}$ 的下标，无需在单调栈中二分查找。

单调栈中的元素出栈时，从哈希表中删除。

这样写，单调栈只需保存 $\textit{nums}[i]$，无需保存下标。

```py [sol-Python3]
class Solution:
    def countConsistent(self, nums: list[int]) -> int:
        data = []
        st = []
        pos = {}
        ans = 0

        for i, x in enumerate(nums):
            # x 入单调栈
            while st and st[-1] >= x:
                del pos[st.pop()]
            st.append(x)
            pos[x] = i

            # 计算右端点为 i 的子数组与值
            for p in data:
                p[0] &= x
            # x 单独一个数作为子数组
            data.append([x, i])

            # 原地去重（相同与值只保留最左边的）
            idx = 1
            for j in range(1, len(data)):
                if data[j][0] != data[j - 1][0]:
                    data[idx] = data[j]
                    idx += 1
            del data[idx:]

            for k, (and_val, left) in enumerate(data):
                # 判断子数组最小值是否为 andVal
                j = pos.get(and_val)
                if j is None:
                    continue

                right = data[k + 1][1] - 1 if k < len(data) - 1 else i
                right = min(right, j)
                ans += max(right - left + 1, 0)

        return ans
```

```java [sol-Java]
class Solution {
    public long countConsistent(int[] nums) {
        List<int[]> data = new ArrayList<>();
        List<Integer> st = new ArrayList<>();
        Map<Integer, Integer> pos = new HashMap<>();
        long ans = 0;

        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];

            // x 入单调栈
            while (!st.isEmpty() && st.getLast() >= x) {
                pos.remove(st.removeLast());
            }
            st.add(x);
            pos.put(x, i);

            // 计算右端点为 i 的子数组与值
            for (int[] p : data) {
                p[0] &= x;
            }
            // x 单独一个数作为子数组
            data.add(new int[]{x, i});

            // 原地去重（相同与值只保留最左边的）
            int m = 1;
            for (int j = 1; j < data.size(); j++) {
                if (data.get(j)[0] != data.get(j - 1)[0]) {
                    data.set(m++, data.get(j));
                }
            }
            data.subList(m, data.size()).clear();

            for (int k = 0; k < m; k++) {
                int andVal = data.get(k)[0];
                // 判断子数组最小值是否为 andVal
                Integer j = pos.get(andVal);
                if (j == null) {
                    continue;
                }

                int left = data.get(k)[1];
                int right = k < m - 1 ? data.get(k + 1)[1] - 1 : i;
                right = Math.min(right, j);
                ans += Math.max(right - left + 1, 0);
            }
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countConsistent(vector<int>& nums) {
        vector<pair<int, int>> data;
        stack<int> st;
        unordered_map<int, int> pos;
        long long ans = 0;

        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];

            // x 入单调栈
            while (!st.empty() && st.top() >= x) {
                pos.erase(st.top());
                st.pop();
            }
            st.push(x);
            pos[x] = i;

            // 计算右端点为 i 的子数组与值
            for (auto& [and_val, _] : data) {
                and_val &= x;
            }
            // x 单独一个数作为子数组
            data.emplace_back(x, i);

            // 原地去重（相同与值只保留最左边的）
            int m = 1;
            for (int j = 1; j < data.size(); j++) {
                if (data[j].first != data[j - 1].first) {
                    data[m++] = data[j];
                }
            }
            data.resize(m);

            for (int k = 0; k < m; k++) {
                auto& [and_val, left] = data[k];
                // 判断子数组最小值是否为 andVal
                auto it = pos.find(and_val);
                if (it == pos.end()) {
                    continue;
                }

                int right = k + 1 < m ? data[k + 1].second - 1 : i;
                right = min(right, it->second);
                ans += max(right - left + 1, 0);
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
func countConsistent(nums []int) (ans int64) {
	type pair struct{ and, left int } // 子数组与值，最小左端点
	data := []pair{}
	st := []int{}
	pos := map[int]int{}

	for i, x := range nums {
		// x 入单调栈
		for len(st) > 0 && st[len(st)-1] >= x {
			delete(pos, st[len(st)-1])
			st = st[:len(st)-1]
		}
		st = append(st, x)
		pos[x] = i

		// 计算右端点为 i 的子数组与值
		for j := range data {
			data[j].and &= x
		}
		// x 单独一个数作为子数组
		data = append(data, pair{x, i})

		// 原地去重（相同与值只保留最左边的）
		idx := 1
		for j := 1; j < len(data); j++ {
			if data[j].and != data[j-1].and {
				data[idx] = data[j]
				idx++
			}
		}
		data = data[:idx]

		for k, p := range data {
			andVal := p.and
			// 判断子数组最小值是否为 andVal
			j, ok := pos[andVal]
			if !ok {
				continue
			}

			right := i
			if k < len(data)-1 {
				right = data[k+1].left - 1
			}
			right = min(right, j)
			ans += int64(max(right-p.left+1, 0))
		}
	}

	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

1. 位运算题单的「**AND/OR LogTrick**」。
2. 单调栈题单。

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
