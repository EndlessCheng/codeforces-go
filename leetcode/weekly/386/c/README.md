[视频讲解](https://www.bilibili.com/video/BV1qx421179t/) 第三题。

## 更形象的题意

题意有点抽象，形象地解释一下：

你有 $n$ 门课程需要考试，第 $i$ 门课程需要用 $\textit{nums}[i]$ 天复习。同一天只能复习一门课程。

在第 $i$ 天，你可以选择参加第 $\textit{changeIndices}[i]$ 门课程的考试。考试这一天不能复习。

搞定所有课程的复习+考试，至少要多少天？

## 方法一：二分答案+正向遍历

#### 提示 1

答案越大，越能够搞定所有课程，反之越不能。

有单调性，可以**二分答案**。

#### 提示 2

考试的时间越晚越好，这样我们能有更充足的时间复习。

设二分的答案为 $\textit{mx}$。在 $\textit{mx}$ 天内，设 $\textit{i}$ 在 $\textit{changeIndices}$ 中出现的最后下标为 $\textit{lastT}[\textit{i}]$，即第 $i$ 门课程的最晚考试时间。如果 $i$ 不在 $\textit{changeIndices}$ 的前 $\textit{mx}$ 个数中，二分返回 `false`。

- 初始化 $\textit{cnt}=0$，遍历 $\textit{changeIndices}$ 的前 $\textit{mx}$ 个数。
- 如果 $i\ne \textit{lastT}[i]$，这一天只能用来复习（或者什么也不做），但还不知道要复习哪一门课程，所以暂时记录一下，把 $\textit{cnt}$ 加一。
- 如果 $i= \textit{lastT}[i]$，先从 $\textit{cnt}$ 中消耗 $\textit{nums}[i]$ 天用来复习，然后考试。如果 $\textit{cnt}<\textit{nums}[i]$ 则无法完成考试，二分返回 `false`。

如果遍历中没有返回 `false`，二分返回 `true`。

下面代码用的开区间二分（其它写法也可以），原理请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

```py [sol-Python3]
class Solution:
    def earliestSecondToMarkIndices(self, nums: List[int], changeIndices: List[int]) -> int:
        n, m = len(nums), len(changeIndices)
        if n > m: return -1

        def check(mx: int) -> bool:
            last_t = [-1] * n
            for t, idx in enumerate(changeIndices[:mx]):
                last_t[idx - 1] = t
            if -1 in last_t:  # 有课程没有考试时间
                return False

            cnt = 0
            for i, idx in enumerate(changeIndices[:mx]):
                idx -= 1
                if i == last_t[idx]:  # 考试
                    if nums[idx] > cnt:  # 没时间复习
                        return False
                    cnt -= nums[idx]  # 复习这门课程
                else:
                    cnt += 1  # 留着后面用
            return True

        left = n + sum(nums)
        ans = left + bisect_left(range(left, m + 1), True, key=check)
        return -1 if ans > m else ans
```

```java [sol-Java]
class Solution {
    public int earliestSecondToMarkIndices(int[] nums, int[] changeIndices) {
        int n = nums.length;
        int m = changeIndices.length;
        if (n > m) {
            return -1;
        }

        int[] lastT = new int[n];
        int left = n - 1, right = m + 1;
        while (left + 1 < right) {
            int mid = (left + right) / 2;
            if (check(nums, changeIndices, lastT, mid)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right > m ? -1 : right;
    }

    private boolean check(int[] nums, int[] changeIndices, int[] lastT, int mx) {
        Arrays.fill(lastT, -1);
        for (int t = 0; t < mx; t++) {
            lastT[changeIndices[t] - 1] = t;
        }
        for (int t : lastT) {
            if (t < 0) { // 有课程没有考试时间
                return false;
            }
        }

        int cnt = 0;
        for (int i = 0; i < mx; i++) {
            int idx = changeIndices[i] - 1;
            if (i == lastT[idx]) { // 考试
                if (nums[idx] > cnt) { // 没时间复习
                    return false;
                }
                cnt -= nums[idx]; // 复习这门课程
            } else {
                cnt++; // 留着后面用
            }
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int earliestSecondToMarkIndices(vector<int> &nums, vector<int> &changeIndices) {
        int n = nums.size(), m = changeIndices.size();
        if (n > m) return -1;

        vector<int> last_t(n);
        auto check = [&](int mx) -> bool {
            ranges::fill(last_t, -1);
            for (int t = 0; t < mx; t++) {
                last_t[changeIndices[t] - 1] = t;
            }
            if (ranges::find(last_t, -1) != last_t.end()) { // 有课程没有考试时间
                return false;
            }

            int cnt = 0;
            for (int i = 0; i < mx; i++) {
                int idx = changeIndices[i] - 1;
                if (i == last_t[idx]) { // 考试
                    if (nums[idx] > cnt) { // 没时间复习
                        return false;
                    }
                    cnt -= nums[idx]; // 复习这门课程
                } else {
                    cnt++; // 留着后面用
                }
            }
            return true;
        };

        int left = n - 1, right = m + 1;
        while (left + 1 < right) {
            int mid = (left + right) / 2;
            (check(mid) ? right : left) = mid;
        }
        return right > m ? -1 : right;
    }
};
```

```go [sol-Go]
func earliestSecondToMarkIndices(nums, changeIndices []int) int {
	n, m := len(nums), len(changeIndices)
	if n > m {
		return -1
	}

	lastT := make([]int, n)
	ans := n + sort.Search(m+1-n, func(mx int) bool {
		mx += n
		clear(lastT)
		for t, idx := range changeIndices[:mx] {
			lastT[idx-1] = t + 1
		}
		if slices.Contains(lastT, 0) { // 有课程没有考试时间
			return false
		}

		cnt := 0
		for i, idx := range changeIndices[:mx] {
			idx--
			if i == lastT[idx]-1 { // 考试
				if nums[idx] > cnt { // 没时间复习
					return false
				}
				cnt -= nums[idx] // 复习这门课程
			} else {
				cnt++ // 留着后面用
			}
		}
		return true
	})
	if ans > m {
		return -1
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m\log m)$，其中 $m$ 为 $\textit{changeIndices}$ 的长度。二分的时候保证 $n\le m$，时间复杂度以 $m$ 为主。
- 空间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。

## 方法二：二分答案+逆向遍历

也可以倒着遍历 $\textit{changeIndices}$ 的前 $\textit{mx}$ 个数。

- 初始化需要考试的课程数 $\textit{exam}=n$ 和需要复习的天数 $\textit{study}=0$。
- 如果第一次遇到 $\textit{changeIndices}[i]$，那么就考试，把 $\textit{exam}$ 减一，把 $\textit{study}$ 增加这门课程需要的复习天数。
- 否则这天用来复习，把 $\textit{study}$ 减一（前提是 $\textit{study}>0$）。

最后如果 $\textit{exam}=0$ 且 $\textit{study}=0$，就表示所有课程都考完了，并且考试前有足够的天数用来复习。

```py [sol-Python3]
class Solution:
    def earliestSecondToMarkIndices(self, nums: List[int], changeIndices: List[int]) -> int:
        n, m = len(nums), len(changeIndices)
        if n > m: return -1

        done = [0] * n  # 避免反复创建和初始化数组
        def check(mx: int) -> bool:
            exam, study = n, 0
            for i in range(mx - 1, -1, -1):
                idx = changeIndices[i] - 1
                if done[idx] != mx:
                    done[idx] = mx
                    exam -= 1  # 考试
                    study += nums[idx]  # 需要复习的天数
                elif study:
                    study -= 1  # 复习
            return exam == 0 and study == 0  # 考完了并且复习完了

        left = n + sum(nums)
        ans = left + bisect_left(range(left, m + 1), True, key=check)
        return -1 if ans > m else ans
```

```java [sol-Java]
class Solution {
    public int earliestSecondToMarkIndices(int[] nums, int[] changeIndices) {
        int n = nums.length;
        int m = changeIndices.length;
        if (n > m) {
            return -1;
        }

        int[] done = new int[n]; // 避免反复创建和初始化数组
        int left = n - 1, right = m + 1;
        while (left + 1 < right) {
            int mid = (left + right) / 2;
            if (check(nums, changeIndices, done, mid)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right > m ? -1 : right;
    }

    private boolean check(int[] nums, int[] changeIndices, int[] done, int mx) {
        int exam = nums.length;
        int study = 0;
        for (int i = mx - 1; i >= 0 && study <= i + 1; i--) { // 要复习的天数不能太多
            int idx = changeIndices[i] - 1;
            if (done[idx] != mx) {
                done[idx] = mx;
                exam--; // 考试
                study += nums[idx]; // 需要复习的天数
            } else if (study > 0) {
                study--; // 复习
            }
        }
        return exam == 0 && study == 0; // 考完了并且复习完了
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int earliestSecondToMarkIndices(vector<int> &nums, vector<int> &changeIndices) {
        int n = nums.size(), m = changeIndices.size();
        if (n > m) return -1;

        vector<int> done(n); // 避免反复创建和初始化数组
        auto check = [&](int mx) -> bool {
            int exam = n, study = 0;
            for (int i = mx - 1; i >= 0 && study <= i + 1; i--) { // 要复习的天数不能太多
                int idx = changeIndices[i] - 1;
                if (done[idx] != mx) {
                    done[idx] = mx;
                    exam--; // 考试
                    study += nums[idx]; // 需要复习的天数
                } else if (study) {
                    study--; // 复习
                }
            }
            return exam == 0 && study == 0; // 考完了并且复习完了
        };

        int left = n - 1, right = m + 1;
        while (left + 1 < right) {
            int mid = (left + right) / 2;
            (check(mid) ? right : left) = mid;
        }
        return right > m ? -1 : right;
    }
};
```

```go [sol-Go]
func earliestSecondToMarkIndices(nums, changeIndices []int) int {
	n, m := len(nums), len(changeIndices)
	if n > m {
		return -1
	}

	done := make([]int, n) // 避免反复创建和初始化数组
	ans := n + sort.Search(m+1-n, func(mx int) bool {
		mx += n
		exam, study := n, 0
		for i := mx - 1; i >= 0; i-- {
			idx := changeIndices[i] - 1
			if done[idx] != mx {
				done[idx] = mx
				exam-- // 考试
				study += nums[idx] // 需要复习的天数
			} else if study > 0 {
				study-- // 复习
			}
		}
		return exam == 0 && study == 0 // 考完了并且复习完了
	})
	if ans > m {
		return -1
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m\log m)$，其中 $m$ 为 $\textit{changeIndices}$ 的长度。二分的时候保证 $n\le m$，时间复杂度以 $m$ 为主。
- 空间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。

#### 题单：二分答案

- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
