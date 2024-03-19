[视频讲解](https://www.bilibili.com/video/BV1zt4y1R7Tc/) 第四题。

## 前置知识

1. [我在知乎上对 KMP 算法的讲解](https://www.zhihu.com/question/21923021/answer/37475572)
2. [二分查找原理](https://www.bilibili.com/video/BV1AP41137w7/)

## 写法一：KMP+二分查找

1. 用 KMP 求出 $a$ 在 $s$ 中的所有出现位置，记作 $\textit{posA}$。
2. 用 KMP 求出 $b$ 在 $s$ 中的所有出现位置，记作 $\textit{posB}$。
3. 遍历 $\textit{posA}$ 中的下标 $i$，在 $\textit{posB}$ 中二分查找离 $i$ 最近的 $j$。如果 $|i-j|\le k$，则把 $i$ 加入答案。

```py [sol-Python3]
class Solution:
    def beautifulIndices(self, s: str, a: str, b: str, k: int) -> List[int]:
        pos_a = self.kmp(s, a)
        pos_b = self.kmp(s, b)

        ans = []
        for i in pos_a:
            bi = bisect_left(pos_b, i)
            if bi < len(pos_b) and pos_b[bi] - i <= k or \
               bi > 0 and i - pos_b[bi - 1] <= k:
                ans.append(i)
        return ans

    def kmp(self, text: str, pattern: str) -> List[int]:
        m = len(pattern)
        pi = [0] * m
        c = 0
        for i in range(1, m):
            v = pattern[i]
            while c and pattern[c] != v:
                c = pi[c - 1]
            if pattern[c] == v:
                c += 1
            pi[i] = c

        res = []
        c = 0
        for i, v in enumerate(text):
            while c and pattern[c] != v:
                c = pi[c - 1]
            if pattern[c] == v:
                c += 1
            if c == len(pattern):
                res.append(i - m + 1)
                c = pi[c - 1]
        return res
```

```java [sol-Java]
class Solution {
    public List<Integer> beautifulIndices(String s, String a, String b, int k) {
        char[] text = s.toCharArray();
        List<Integer> posA = kmp(text, a.toCharArray());
        List<Integer> posB = kmp(text, b.toCharArray());

        List<Integer> ans = new ArrayList<>();
        for (int i : posA) {
            int bi = lowerBound(posB, i);
            if (bi < posB.size() && posB.get(bi) - i <= k ||
                bi > 0 && i - posB.get(bi - 1) <= k) {
                ans.add(i);
            }
        }
        return ans;
    }

    private List<Integer> kmp(char[] text, char[] pattern) {
        int m = pattern.length;
        int[] pi = new int[m];
        int c = 0;
        for (int i = 1; i < m; i++) {
            char v = pattern[i];
            while (c > 0 && pattern[c] != v) {
                c = pi[c - 1];
            }
            if (pattern[c] == v) {
                c++;
            }
            pi[i] = c;
        }

        List<Integer> res = new ArrayList<>();
        c = 0;
        for (int i = 0; i < text.length; i++) {
            char v = text[i];
            while (c > 0 && pattern[c] != v) {
                c = pi[c - 1];
            }
            if (pattern[c] == v) {
                c++;
            }
            if (c == m) {
                res.add(i - m + 1);
                c = pi[c - 1];
            }
        }
        return res;
    }

    // 开区间写法
    // 请看 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(List<Integer> nums, int target) {
        int left = -1, right = nums.size(); // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[left] < target
            // nums[right] >= target
            int mid = (left + right) >>> 1;
            if (nums.get(mid) < target) {
                left = mid; // 范围缩小到 (mid, right)
            } else {
                right = mid; // 范围缩小到 (left, mid)
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> beautifulIndices(string s, string a, string b, int k) {
        vector<int> posA = kmp(s, a);
        vector<int> posB = kmp(s, b);

        vector<int> ans;
        for (int i: posA) {
            auto it = lower_bound(posB.begin(), posB.end(), i);
            if (it != posB.end() && *it - i <= k ||
                it != posB.begin() && i - *--it <= k) {
                ans.push_back(i);
            }
        }
        return ans;
    }

private:
    vector<int> kmp(string &text, string &pattern) {
        int m = pattern.length();
        vector<int> pi(m);
        int c = 0;
        for (int i = 1; i < m; i++) {
            char v = pattern[i];
            while (c && pattern[c] != v) {
                c = pi[c - 1];
            }
            if (pattern[c] == v) {
                c++;
            }
            pi[i] = c;
        }

        vector<int> res;
        c = 0;
        for (int i = 0; i < text.length(); i++) {
            char v = text[i];
            while (c && pattern[c] != v) {
                c = pi[c - 1];
            }
            if (pattern[c] == v) {
                c++;
            }
            if (c == m) {
                res.push_back(i - m + 1);
                c = pi[c - 1];
            }
        }
        return res;
    }
};
```

```go [sol-Go]
func beautifulIndices(s, a, b string, k int) (ans []int) {
	posA := kmp(s, a)
	posB := kmp(s, b)

	for _, i := range posA {
		bi := sort.SearchInts(posB, i)
		// 离 i 最近的 j 是 posB[bi] 或 posB[bi-1]
		if bi < len(posB) && posB[bi]-i <= k || bi > 0 && i-posB[bi-1] <= k {
			ans = append(ans, i)
		}
	}
	return
}

func kmp(text, pattern string) (pos []int) {
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
	for i, v := range text {
		for cnt > 0 && pattern[cnt] != byte(v) {
			cnt = pi[cnt-1]
		}
		if pattern[cnt] == byte(v) {
			cnt++
		}
		if cnt == m {
			pos = append(pos, i-m+1)
			cnt = pi[cnt-1]
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 写法二：KMP+双指针

由于 $\textit{posA}$ 和 $\textit{posB}$ 都是有序的，也可以用双指针做。

```py [sol-Python3]
class Solution:
    def beautifulIndices(self, s: str, a: str, b: str, k: int) -> List[int]:
        pos_a = self.kmp(s, a)
        pos_b = self.kmp(s, b)

        ans = []
        j, m = 0, len(pos_b)
        for i in pos_a:
            while j < m and pos_b[j] < i - k:
                j += 1
            if j < m and pos_b[j] <= i + k:
                ans.append(i)
        return ans

    def kmp(self, text: str, pattern: str) -> List[int]:
        m = len(pattern)
        pi = [0] * m
        c = 0
        for i in range(1, m):
            v = pattern[i]
            while c and pattern[c] != v:
                c = pi[c - 1]
            if pattern[c] == v:
                c += 1
            pi[i] = c

        res = []
        c = 0
        for i, v in enumerate(text):
            while c and pattern[c] != v:
                c = pi[c - 1]
            if pattern[c] == v:
                c += 1
            if c == len(pattern):
                res.append(i - m + 1)
                c = pi[c - 1]
        return res
```

```java [sol-Java]
class Solution {
    public List<Integer> beautifulIndices(String s, String a, String b, int k) {
        char[] text = s.toCharArray();
        List<Integer> posA = kmp(text, a.toCharArray());
        List<Integer> posB = kmp(text, b.toCharArray());

        List<Integer> ans = new ArrayList<>();
        int j = 0, m = posB.size();
        for (int i : posA) {
            while (j < m && posB.get(j) < i - k) {
                j++;
            }
            if (j < m && posB.get(j) <= i + k) {
                ans.add(i);
            }
        }
        return ans;
    }

    private List<Integer> kmp(char[] text, char[] pattern) {
        int m = pattern.length;
        int[] pi = new int[m];
        int c = 0;
        for (int i = 1; i < m; i++) {
            char v = pattern[i];
            while (c > 0 && pattern[c] != v) {
                c = pi[c - 1];
            }
            if (pattern[c] == v) {
                c++;
            }
            pi[i] = c;
        }

        List<Integer> res = new ArrayList<>();
        c = 0;
        for (int i = 0; i < text.length; i++) {
            char v = text[i];
            while (c > 0 && pattern[c] != v) {
                c = pi[c - 1];
            }
            if (pattern[c] == v) {
                c++;
            }
            if (c == m) {
                res.add(i - m + 1);
                c = pi[c - 1];
            }
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> beautifulIndices(string s, string a, string b, int k) {
        vector<int> posA = kmp(s, a);
        vector<int> posB = kmp(s, b);

        vector<int> ans;
        int j = 0, m = posB.size();
        for (int i : posA) {
            while (j < m && posB[j] < i - k) {
                j++;
            }
            if (j < m && posB[j] <= i + k) {
                ans.push_back(i);
            }
        }
        return ans;
    }

private:
    vector<int> kmp(string &text, string &pattern) {
        int m = pattern.length();
        vector<int> pi(m);
        int c = 0;
        for (int i = 1; i < m; i++) {
            char v = pattern[i];
            while (c && pattern[c] != v) {
                c = pi[c - 1];
            }
            if (pattern[c] == v) {
                c++;
            }
            pi[i] = c;
        }

        vector<int> res;
        c = 0;
        for (int i = 0; i < text.length(); i++) {
            char v = text[i];
            while (c && pattern[c] != v) {
                c = pi[c - 1];
            }
            if (pattern[c] == v) {
                c++;
            }
            if (c == m) {
                res.push_back(i - m + 1);
                c = pi[c - 1];
            }
        }
        return res;
    }
};
```

```go [sol-Go]
func beautifulIndices(s, a, b string, k int) (ans []int) {
	posA := kmp(s, a)
	posB := kmp(s, b)

	j, m := 0, len(posB)
	for _, i := range posA {
		for j < m && posB[j] < i-k {
			j++
		}
		if j < m && posB[j] <= i+k {
			ans = append(ans, i)
		}
	}
	return
}

func kmp(text, pattern string) (pos []int) {
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
	for i, v := range text {
		for cnt > 0 && pattern[cnt] != byte(v) {
			cnt = pi[cnt-1]
		}
		if pattern[cnt] == byte(v) {
			cnt++
		}
		if cnt == m {
			pos = append(pos, i-m+1)
			cnt = pi[cnt-1]
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。注意二重循环中 `j++` 的次数不会超过 $n$，所以二重循环的时间复杂度是 $\mathcal{O}(n)$ 的。
- 空间复杂度：$\mathcal{O}(n)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
