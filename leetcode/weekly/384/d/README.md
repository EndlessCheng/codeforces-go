把 $\textit{nums}$ 的相邻元素，根据题目规定的大小关系，转换成 $1,0,-1$，得到一个长为 $n-1$ 的数组 $b$。

问题相当于问 $b$ 中有多少个连续子数组等于 $\textit{pattern}$。

这是一个标准的字符串匹配问题（本题匹配的是数字不是字符），可以用 KMP 或者 Z 函数解决。

## 方法一：KMP

关于 KMP 的原理，请看我在知乎上的 [这篇讲解](https://www.zhihu.com/question/21923021/answer/37475572)。

```py [sol-Python3]
class Solution:
    def countMatchingSubarrays(self, nums: List[int], pattern: List[int]) -> int:
        m = len(pattern)
        pi = [0] * m
        cnt = 0
        for i in range(1, m):
            v = pattern[i]
            while cnt and pattern[cnt] != v:
                cnt = pi[cnt - 1]
            if pattern[cnt] == v:
                cnt += 1
            pi[i] = cnt

        ans = cnt = 0
        for x, y in pairwise(nums):
            v = (y > x) - (y < x)
            while cnt and pattern[cnt] != v:
                cnt = pi[cnt - 1]
            if pattern[cnt] == v:
                cnt += 1
            if cnt == m:
                ans += 1
                cnt = pi[cnt - 1]
        return ans
```

```java [sol-Java]
class Solution {
    public int countMatchingSubarrays(int[] nums, int[] pattern) {
        int m = pattern.length;
        int[] pi = new int[m];
        int cnt = 0;
        for (int i = 1; i < m; i++) {
            int v = pattern[i];
            while (cnt > 0 && pattern[cnt] != v) {
                cnt = pi[cnt - 1];
            }
            if (pattern[cnt] == v) {
                cnt++;
            }
            pi[i] = cnt;
        }

        int ans = 0;
        cnt = 0;
        for (int i = 1; i < nums.length; i++) {
            int v = Integer.compare(nums[i], nums[i - 1]);
            while (cnt > 0 && pattern[cnt] != v) {
                cnt = pi[cnt - 1];
            }
            if (pattern[cnt] == v) {
                cnt++;
            }
            if (cnt == m) {
                ans++;
                cnt = pi[cnt - 1];
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countMatchingSubarrays(vector<int> &nums, vector<int> &pattern) {
        int m = pattern.size();
        vector<int> pi(m);
        int cnt = 0;
        for (int i = 1; i < m; i++) {
            int v = pattern[i];
            while (cnt && pattern[cnt] != v) {
                cnt = pi[cnt - 1];
            }
            cnt += pattern[cnt] == v;
            pi[i] = cnt;
        }

        int ans = 0;
        cnt = 0;
        for (int i = 1; i < nums.size(); i++) {
            int x = nums[i - 1], y = nums[i];
            int v = (y > x) - (y < x);
            while (cnt && pattern[cnt] != v) {
                cnt = pi[cnt - 1];
            }
            cnt += pattern[cnt] == v;
            if (cnt == m) {
                ans++;
                cnt = pi[cnt - 1];
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func countMatchingSubarrays(nums, pattern []int) (ans int) {
	m := len(pattern)
	pi := make([]int, m)
	cnt := 0
	for i := 1; i < m; i++ {
		v := pattern[i]
		for cnt > 0 && pattern[cnt] != v {
			cnt = pi[cnt-1]
		}
		if pattern[cnt] == v {
			cnt++
		}
		pi[i] = cnt
	}

	cnt = 0
	for i := 1; i < len(nums); i++ {
		v := cmp.Compare(nums[i], nums[i-1])
		for cnt > 0 && pattern[cnt] != v {
			cnt = pi[cnt-1]
		}
		if pattern[cnt] == v {
			cnt++
		}
		if cnt == m {
			ans++
			cnt = pi[cnt-1]
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(m)$，其中 $m$ 为 $\textit{pattern}$ 的长度。

## 方法二：Z 函数（扩展 KMP）

上场周赛讲了 [Z 函数的原理](https://www.bilibili.com/video/BV1it421W7D8/)。

我们可以把 $\textit{pattern}$ 拼在 $b$ 的前面（为了防止匹配越界，中间插入一个不在数组中的数字，比如 $2$），根据 Z 函数的定义，只要 $z[i] = m$，我们就找到了一个匹配。

```py [sol-Python3]
class Solution:
    def countMatchingSubarrays(self, nums: List[int], pattern: List[int]) -> int:
        m = len(pattern)
        pattern.append(2)
        pattern.extend((y > x) - (y < x) for x, y in pairwise(nums))

        n = len(pattern)
        z = [0] * n
        l = r = 0
        for i in range(1, n):
            if i <= r:
                z[i] = min(z[i - l], r - i + 1)
            while i + z[i] < n and pattern[z[i]] == pattern[i + z[i]]:
                l, r = i, i + z[i]
                z[i] += 1

        return sum(lcp == m for lcp in z[m + 1:])
```

```java [sol-Java]
class Solution {
    public int countMatchingSubarrays(int[] nums, int[] pattern) {
        int m = pattern.length;
        int[] s = Arrays.copyOf(pattern, m + nums.length);
        s[m] = 2;
        for (int i = 1; i < nums.length; i++) {
            s[m + i] = Integer.compare(nums[i], nums[i - 1]);
        }

        int n = s.length;
        int[] z = new int[n];
        int l = 0, r = 0;
        for (int i = 1; i < n; i++) {
            if (i <= r) {
                z[i] = Math.min(z[i - l], r - i + 1);
            }
            while (i + z[i] < n && s[z[i]] == s[i + z[i]]) {
                l = i;
                r = i + z[i];
                z[i]++;
            }
        }

        int ans = 0;
        for (int i = m + 1; i < n; i++) {
            if (z[i] == m) {
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countMatchingSubarrays(vector<int> &nums, vector<int> &pattern) {
        int m = pattern.size();
        pattern.push_back(2);
        for (int i = 1; i < nums.size(); i++) {
            int x = nums[i - 1], y = nums[i];
            pattern.push_back((y > x) - (y < x));
        }

        int n = pattern.size();
        vector<int> z(n);
        int l = 0, r = 0;
        for (int i = 1; i < n; i++) {
            if (i <= r) {
                z[i] = min(z[i - l], r - i + 1);
            }
            while (i + z[i] < n && pattern[z[i]] == pattern[i + z[i]]) {
                l = i;
                r = i + z[i];
                z[i]++;
            }
        }

        int ans = 0;
        for (int i = m + 1; i < n; i++) {
            ans += z[i] == m;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countMatchingSubarrays(nums, pattern []int) (ans int) {
	m := len(pattern)
	pattern = append(pattern, 2)
	for i := 1; i < len(nums); i++ {
		pattern = append(pattern, cmp.Compare(nums[i], nums[i-1]))
	}

	n := len(pattern)
	z := make([]int, n)
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min(z[i-l], r-i+1)
		}
		for i+z[i] < n && pattern[z[i]] == pattern[i+z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
	}

	for _, lcp := range z[m+1:] {
		if lcp == m {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
