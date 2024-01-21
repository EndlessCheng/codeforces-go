[视频讲解](https://www.bilibili.com/video/BV1oV411D7gB/) 第四题。

第一段的第一个数是确定的，即 $\textit{nums}[0]$。

如果知道了第二段的第一个数的位置（记作 $p$），第三段的第一个数的位置，……，第 $k$ 段的第一个数的位置（记作 $q$），那么这个划分方案也就确定了。

考虑到 $q-p \le \textit{dist}$，本题相当于在一个大小固定为 $\textit{dist}+1$ 的滑动窗口内，求前 $k-1$ 小的元素和。

仿照 [480. 滑动窗口中位数](https://leetcode.cn/problems/sliding-window-median/)，这可以用两个有序集合来做：

1. 初始化两个有序集合 $L$ 和 $R$。注意：为方便计算，把 $k$ 减一。
2. 把 $\textit{nums}[1]$ 到 $\textit{nums}[\textit{dist}+1]$ 加到 $L$ 中。
3. 保留 $L$ 最小的 $k$ 个数，把其余数丢到 $R$ 中。
4. 从 $i=\textit{dist}+2$ 开始滑窗。
5. 先把 $\textit{out} = \textit{nums}[i-\textit{dist}-1]$ 移出窗口：如果 $\textit{out}$ 在 $L$ 中，就从 $L$ 中移除，否则从 $R$ 中移除。
6. 然后把 $\textit{in} = \textit{nums}[i]$ 移入窗口：如果 $\textit{in}$ 小于 $L$ 中的最大元素，则加入 $L$，否则加入 $R$。
7. 上面两步做完后，如果 $L$ 中的元素个数小于 $k$（等于 $k-1$），则从 $R$ 中取一个最小元素加入 $L$；反之，如果 $L$ 中的元素个数大于 $k$（等于 $k+1$），则从 $L$ 中取一个最大元素加入 $R$。

上述过程维护 $L$ 中元素之和 $\textit{sumL}$，取 $\textit{sumL}$ 的最小值，即为答案。

```py [sol-Python3]
from sortedcontainers import SortedList

class Solution:
    def minimumCost(self, nums: List[int], k: int, dist: int) -> int:
        k -= 1
        sum_left = sum(nums[:dist + 2])
        L = SortedList(nums[1:dist + 2])
        R = SortedList()

        def L2R() -> None:
            x = L.pop()
            nonlocal sum_left
            sum_left -= x
            R.add(x)

        def R2L() -> None:
            x = R.pop(0)
            nonlocal sum_left
            sum_left += x
            L.add(x)

        while len(L) > k:
            L2R()

        ans = sum_left
        for i in range(dist + 2, len(nums)):
            # 移除 out
            out = nums[i - dist - 1]
            if out in L:
                sum_left -= out
                L.remove(out)
            else:
                R.remove(out)

            # 添加 in
            in_val = nums[i]
            if in_val < L[-1]:
                sum_left += in_val
                L.add(in_val)
            else:
                R.add(in_val)

            # 维护大小
            if len(L) == k - 1:
                R2L()
            elif len(L) == k + 1:
                L2R()

            ans = min(ans, sum_left)

        return ans
```

```cpp [sol-C++]
class Solution {
public:
    long long minimumCost(vector<int> &nums, int k, int dist) {
        k--;
        long long sum = accumulate(nums.begin(), nums.begin() + dist + 2, 0LL);
        multiset<int> L(nums.begin() + 1, nums.begin() + dist + 2), R;
        auto L2R = [&]() {
            int x = *L.rbegin();
            sum -= x;
            L.erase(L.find(x));
            R.insert(x);
        };
        auto R2L = [&]() {
            int x = *R.begin();
            sum += x;
            R.erase(R.find(x));
            L.insert(x);
        };
        while (L.size() > k) {
            L2R();
        }

        long long ans = sum;
        for (int i = dist + 2; i < nums.size(); i++) {
            // 移除 out
            int out = nums[i - dist - 1];
            auto it = L.find(out);
            if (it != L.end()) {
                sum -= out;
                L.erase(it);
            } else {
                R.erase(R.find(out));
            }

            // 添加 in
            int in = nums[i];
            if (in < *L.rbegin()) {
                sum += in;
                L.insert(in);
            } else {
                R.insert(in);
            }

            // 维护大小
            if (L.size() == k - 1) {
                R2L();
            } else if (L.size() == k + 1) {
                L2R();
            }

            ans = min(ans, sum);
        }
        return ans;
    }
};
```

```java [sol-Java]
public class Solution {
    public long minimumCost(int[] nums, int k, int dist) {
        k--;
        sumL = nums[0];
        for (int i = 1; i < dist + 2; i++) {
            sumL += nums[i];
            L.merge(nums[i], 1, Integer::sum);
        }
        sizeL = dist + 1;
        while (sizeL > k) {
            l2r();
        }

        long ans = sumL;
        for (int i = dist + 2; i < nums.length; i++) {
            // 移除 out
            int out = nums[i - dist - 1];
            if (L.containsKey(out)) {
                sumL -= out;
                sizeL--;
                removeOne(L, out);
            } else {
                removeOne(R, out);
            }

            // 添加 in
            int in = nums[i];
            if (in < L.lastKey()) {
                sumL += in;
                sizeL++;
                L.merge(in, 1, Integer::sum);
            } else {
                R.merge(in, 1, Integer::sum);
            }

            // 维护大小
            if (sizeL == k - 1) {
                r2l();
            } else if (sizeL == k + 1) {
                l2r();
            }

            ans = Math.min(ans, sumL);
        }
        return ans;
    }

    private long sumL;
    private int sizeL;
    private final TreeMap<Integer, Integer> L = new TreeMap<>();
    private final TreeMap<Integer, Integer> R = new TreeMap<>();

    private void l2r() {
        int x = L.lastKey();
        removeOne(L, x);
        sumL -= x;
        sizeL--;
        R.merge(x, 1, Integer::sum);
    }

    private void r2l() {
        int x = R.firstKey();
        removeOne(R, x);
        sumL += x;
        sizeL++;
        L.merge(x, 1, Integer::sum);
    }

    private void removeOne(Map<Integer, Integer> m, int x) {
        int cnt = m.get(x);
        if (cnt > 1) {
            m.put(x, cnt - 1);
        } else {
            m.remove(x);
        }
    }
}
```

```go [sol-Go]
import "github.com/emirpasic/gods/trees/redblacktree"

func minimumCost(nums []int, k, dist int) int64 {
	k--
	L := redblacktree.NewWithIntComparator()
	R := redblacktree.NewWithIntComparator()
	add := func(t *redblacktree.Tree, x int) {
		c, ok := t.Get(x)
		if ok {
			t.Put(x, c.(int)+1)
		} else {
			t.Put(x, 1)
		}
	}
	del := func(t *redblacktree.Tree, x int) {
		c, _ := t.Get(x)
		if c.(int) > 1 {
			t.Put(x, c.(int)-1)
		} else {
			t.Remove(x)
		}
	}

	sumL := nums[0]
	for _, x := range nums[1 : dist+2] {
		sumL += x
		add(L, x)
	}
	sizeL := dist + 1

	l2r := func() {
		x := L.Right().Key.(int)
		sumL -= x
		sizeL--
		del(L, x)
		add(R, x)
	}
	r2l := func() {
		x := R.Left().Key.(int)
		sumL += x
		sizeL++
		del(R, x)
		add(L, x)
	}
	for sizeL > k {
		l2r()
	}

	ans := sumL
	for i := dist + 2; i < len(nums); i++ {
		// 移除 out
		out := nums[i-dist-1]
		if _, ok := L.Get(out); ok {
			sumL -= out
			sizeL--
			del(L, out)
		} else {
			del(R, out)
		}

		// 添加 in
		in := nums[i]
		if in < L.Right().Key.(int) {
			sumL += in
			sizeL++
			add(L, in)
		} else {
			add(R, in)
		}

		// 维护大小
		if sizeL == k-1 {
			r2l()
		} else if sizeL == k+1 {
			l2r()
		}

		ans = min(ans, sumL)
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log \textit{dist})$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(\textit{dist})$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
