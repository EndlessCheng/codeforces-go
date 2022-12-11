[视频讲解](https://www.bilibili.com/video/BV1kR4y1r7Df/) 已出炉，欢迎点赞三连，在评论区分享你对这场双周赛的看法~

---

统计满足 $x=\textit{nums}_1[i]=\textit{nums}_2[i]$ 的数对的个数 $\textit{swapCnt}$，以及 $x$ 的众数 $\textit{mode}$ 及其出现次数 $\textit{modeCnt}$。

分类讨论：
- 如果 $\textit{modeCnt}$ 没有超过 $\textit{swapCnt}$ 的一半：
  - 如果 $\textit{swapCnt}$ 是偶数，那么两两交换即可；
  - 如果 $\textit{swapCnt}$ 是奇数，那么至少有三种不同的 $x$，其中一个数必然可以和 $\textit{nums}_1[0]$ 交换；
  - 因此这种情况下，代价就是这些 $x$ 的下标之和。
- 如果 $\textit{modeCnt}$ 超过 $\textit{swapCnt}$ 的一半，或者说 $\textit{modeCnt}\cdot 2 > \textit{swapCnt}$，根据鸽巢原理，无法通过重排这些数字，让数组不相等（因为还存在一些 $\textit{mode}$ 仍然相同）。这种情况必须不断寻找其他的满足 $\textit{nums}_1[j]\ne\textit{nums}_2[j]$ 的数对，且数对中的数都不等于 $\textit{mode}$，直到 $\textit{modeCnt}\cdot 2 \le \textit{swapCnt}$ 为止。为了让答案尽量小，应从左到右遍历数组。如果仍然无法满足要求，则返回 $-1$。

```py [sol1-Python3]
class Solution:
    def minimumTotalCost(self, nums1: List[int], nums2: List[int]) -> int:
        ans = swap_cnt = mode_cnt = mode = 0
        cnt = [0] * (len(nums1) + 1)
        for i, (x, y) in enumerate(zip(nums1, nums2)):
            if x == y:
                ans += i
                swap_cnt += 1
                cnt[x] += 1
                if cnt[x] > mode_cnt:
                    mode_cnt, mode = cnt[x], x

        for i, (x, y) in enumerate(zip(nums1, nums2)):
            if mode_cnt * 2 <= swap_cnt: break
            if x != y and x != mode and y != mode:
                ans += i
                swap_cnt += 1
        return ans if mode_cnt * 2 <= swap_cnt else -1
```

```java [sol1-Java]
class Solution {
    public long minimumTotalCost(int[] nums1, int[] nums2) {
        long ans = 0L;
        int swapCnt = 0, modeCnt = 0, mode = 0, n = nums1.length;
        int[] cnt = new int[n + 1];
        for (int i = 0; i < n; ++i) {
            int x = nums1[i];
            if (x == nums2[i]) {
                ans += i;
                ++swapCnt;
                ++cnt[x];
                if (cnt[x] > modeCnt) {
                    modeCnt = cnt[x];
                    mode = x;
                }
            }
        }

        for (int i = 0; i < n && modeCnt * 2 > swapCnt; ++i) {
            int x = nums1[i], y = nums2[i];
            if (x != y && x != mode && y != mode) {
                ans += i;
                ++swapCnt;
            }
        }
        return modeCnt * 2 > swapCnt ? -1 : ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    long long minimumTotalCost(vector<int> &nums1, vector<int> &nums2) {
        long ans = 0L;
        int swap_cnt = 0, mode_cnt = 0, mode, n = nums1.size(), cnt[n + 1];
        memset(cnt, 0, sizeof(cnt));
        for (int i = 0; i < n; ++i)
            if (int x = nums1[i]; x == nums2[i]) {
                ans += i;
                ++swap_cnt;
                ++cnt[x];
                if (cnt[x] > mode_cnt) {
                    mode_cnt = cnt[x];
                    mode = x;
                }
            }

        for (int i = 0; i < n && mode_cnt * 2 > swap_cnt; ++i) {
            int x = nums1[i], y = nums2[i];
            if (x != y && x != mode && y != mode) {
                ans += i;
                ++swap_cnt;
            }
        }
        return mode_cnt * 2 > swap_cnt ? -1 : ans;
    }
};
```

```go [sol1-Go]
func minimumTotalCost(nums1, nums2 []int) (ans int64) {
	var swapCnt, modeCnt, mode int
	cnt := make([]int, len(nums1)+1)
	for i, x := range nums1 {
		if x == nums2[i] {
			ans += int64(i)
			swapCnt++
			cnt[x]++
			if cnt[x] > modeCnt {
				modeCnt, mode = cnt[x], x
			}
		}
	}

	for i, x := range nums1 {
		if modeCnt*2 <= swapCnt {
			break
		}
		if x != nums2[i] && x != mode && nums2[i] != mode {
			ans += int64(i)
			swapCnt++
		}
	}
	if modeCnt*2 > swapCnt {
		return -1
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}_1$ 的长度。
- 空间复杂度：$O(n)$。还可以用**摩尔投票算法**做到 $O(1)$，见 [169. 多数元素](https://leetcode.cn/problems/majority-element/)。
