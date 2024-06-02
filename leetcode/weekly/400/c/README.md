**核心思路**：由于要去掉最小的字母，为了让字典序尽量小，相比去掉前面的字母，去掉后面的字母更好。[视频讲解](https://www.bilibili.com/video/BV1Qx4y1E7zj/) 第三题简单证明了这点。

从左到右遍历 $s$，用 $26$ 个栈模拟。

第 $i$ 个栈维护第 $i$ 个小写字母的下标。

遇到 `*` 时，弹出第一个非空栈的栈顶下标。

最后把所有栈顶下标对应的字母组合起来，即为答案。

## 写法一

```py [sol-Python3]
class Solution:
    def clearStars(self, s: str) -> str:
        st = [[] for _ in range(26)]
        for i, c in enumerate(s):
            if c != '*':
                st[ord(c) - ord('a')].append(i)
                continue
            for p in st:
                if p:
                    p.pop()
                    break
        return ''.join(s[i] for i in sorted(chain.from_iterable(st)))
```

```java [sol-Java]
class Solution {
    public String clearStars(String S) {
        char[] s = S.toCharArray();
        List<Integer>[] st = new ArrayList[26];
        Arrays.setAll(st, i -> new ArrayList<>());
        for (int i = 0; i < s.length; i++) {
            if (s[i] != '*') {
                st[s[i] - 'a'].add(i);
                continue;
            }
            for (List<Integer> p : st) {
                if (!p.isEmpty()) {
                    p.remove(p.size() - 1);
                    break;
                }
            }
        }

        List<Integer> idx = new ArrayList<>();
        for (List<Integer> p : st) {
            idx.addAll(p);
        }
        Collections.sort(idx);

        StringBuilder t = new StringBuilder(idx.size());
        for (int i : idx) {
            t.append(s[i]);
        }
        return t.toString();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string clearStars(string s) {
        vector<int> st[26];
        for (int i = 0; i < s.size(); i++) {
            if (s[i] != '*') {
                st[s[i] - 'a'].push_back(i);
                continue;
            }
            for (auto& p : st) {
                if (!p.empty()) {
                    p.pop_back();
                    break;
                }
            }
        }

        vector<int> idx;
        for (auto& p : st) {
            idx.insert(idx.end(), p.begin(), p.end());
        }
        ranges::sort(idx);

        string t(idx.size(), 0);
        for (int i = 0; i < idx.size(); i++) {
            t[i] = s[idx[i]];
        }
        return t;
    }
};
```

```go [sol-Go]
func clearStars(s string) string {
	st := [26][]int{}
	for i, c := range s {
		if c != '*' {
			st[c-'a'] = append(st[c-'a'], i)
			continue
		}
		for j, p := range st {
			if len(p) > 0 {
				st[j] = p[:len(p)-1]
				break
			}
		}
	}

	idx := []int{}
	for _, p := range st {
		idx = append(idx, p...)
	}
	slices.Sort(idx)

	t := make([]byte, len(idx))
	for i, j := range idx {
		t[i] = s[j]
	}
	return string(t)
}
```

## 写法二

也可以用一个布尔数组标记需要删除的下标。

```py [sol-Python3]
class Solution:
    def clearStars(self, s: str) -> str:
        delete = [False] * len(s)
        st = [[] for _ in range(26)]
        for i, c in enumerate(s):
            if c != '*':
                st[ord(c) - ord('a')].append(i)
                continue
            for p in st:
                if p:
                    delete[p.pop()] = True
                    break
        return ''.join(c for d, c in zip(delete, s) if not d and c != '*')
```

```java [sol-Java]
class Solution {
    public String clearStars(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        boolean[] del = new boolean[n];
        List<Integer>[] st = new ArrayList[26];
        Arrays.setAll(st, i -> new ArrayList<>());
        for (int i = 0; i < n; i++) {
            if (s[i] != '*') {
                st[s[i] - 'a'].add(i);
                continue;
            }
            for (List<Integer> p : st) {
                if (!p.isEmpty()) {
                    del[p.remove(p.size() - 1)] = true;
                    break;
                }
            }
        }

        StringBuilder t = new StringBuilder();
        for (int i = 0; i < n; i++) {
            if (!del[i] && s[i] != '*') {
                t.append(s[i]);
            }
        }
        return t.toString();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string clearStars(string s) {
        int n = s.length();
        vector<int> del(n);
        stack<int> st[26];
        for (int i = 0; i < n; i++) {
            if (s[i] != '*') {
                st[s[i] - 'a'].push(i);
                continue;
            }
            for (auto& p : st) {
                if (!p.empty()) {
                    del[p.top()] = true;
                    p.pop();
                    break;
                }
            }
        }

        string t;
        for (int i = 0; i < n; i++) {
            if (!del[i] && s[i] != '*') {
                t += s[i];
            }
        }
        return t;
    }
};
```

```go [sol-Go]
func clearStars(s string) string {
	del := make([]bool, len(s))
	st := [26][]int{}
	for i, c := range s {
		if c != '*' {
			st[c-'a'] = append(st[c-'a'], i)
			continue
		}
		for j, ps := range st {
			if m := len(ps); m > 0 {
				del[ps[m-1]] = true
				st[j] = ps[:m-1]
				break
			}
		}
	}

	t := []byte{}
	for i, d := range del {
		if !d && s[i] != '*' {
			t = append(t, s[i])
		}
	}
	return string(t)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|$ 为字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(n+|\Sigma|)$。

注：用最小堆或者有序集合，可以把 $\mathcal{O}(n|\Sigma|)$ 优化至 $\mathcal{O}(n\log |\Sigma|)$。

## 思考题

改成删除最大的字母，要怎么做？需要把栈换成哪个数据结构？

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
