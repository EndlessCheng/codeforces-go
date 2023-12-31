[本题视频讲解](https://www.bilibili.com/video/BV1XG411B7bX/)

按照相同字母分组，每组统计相同字母连续出现的长度。例如 aaaabbbabb 把 a 分成一组，组内有长度 $4$ 和长度 $1$；把 b 分成一组，组内有长度 $3$ 和长度 $2$。

单独考虑每一组，按照长度从大到小排序，设长度列表为 $a$。

分类讨论：

- 从最长的特殊子串（$a[0]$）中取三个长度均为 $a[0]-2$ 的特殊子串。例如示例 1 的 aaaa 可以取三个 aa。
- 从最长和次长的特殊子串（$a[0],a[1]$）中取三个长度一样的特殊子串：
  - 如果 $a[0]=a[1]$，那么可以取三个长度均为 $a[0]-1$ 的特殊子串。
  - 如果 $a[0]>a[1]$，那么可以取三个长度均为 $a[1]$ 的特殊子串：从最长中取两个，从次长中取一个。
  - 这两种情况合并成 $\min(a[0]-1, a[1])$。
- 从最长、次长、第三长的的特殊子串（$a[0],a[1],a[2]$）中各取一个长为 $a[2]$ 的特殊子串。

这三种情况取最大值，即

$$
\max(a[0]-2, \min(a[0]-1, a[1]), a[2])
$$

取每一组的最大值，即为答案。

如果答案是 $0$，返回 $-1$。

代码实现时，无需特判 $a$ 数组长度小于 $3$ 的情况，我们只需要在数组末尾加两个 $0$ 即可。

```py [sol-Python3]
class Solution:
    def maximumLength(self, s: str) -> int:
        groups = defaultdict(list)
        cnt = 0
        for i, ch in enumerate(s):
            cnt += 1
            if i == len(s) - 1 or ch != s[i + 1]:
                groups[ch].append(cnt)  # 统计连续字符长度
                cnt = 0

        ans = 0
        for a in groups.values():
            a.sort(reverse=True)
            a.extend([0, 0])  # 假设还有两个空串
            ans = max(ans, a[0] - 2, min(a[0] - 1, a[1]), a[2])
        return ans if ans else -1
```

```java [sol-Java]
public class Solution {
    public int maximumLength(String S) {
        char[] s = S.toCharArray();
        List<Integer>[] groups = new ArrayList[26];
        Arrays.setAll(groups, i -> new ArrayList<>());
        int cnt = 0;
        for (int i = 0; i < s.length; i++) {
            cnt++;
            if (i == s.length - 1 || s[i] != s[i + 1]) {
                groups[s[i] - 'a'].add(cnt); // 统计连续字符长度
                cnt = 0;
            }
        }

        int ans = 0;
        for (List<Integer> a : groups) {
            if (a.isEmpty()) continue;
            a.sort(Collections.reverseOrder());
            a.add(0);
            a.add(0); // 假设还有两个空串
            ans = Math.max(ans, Math.max(a.get(0) - 2, Math.max(Math.min(a.get(0) - 1, a.get(1)), a.get(2))));
        }
        return ans > 0 ? ans : -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumLength(string s) {
        vector<int> groups[26];
        int cnt = 0, n = s.length();
        for (int i = 0; i < n; i++) {
            cnt++;
            if (i + 1 == n || s[i] != s[i + 1]) {
                groups[s[i] - 'a'].push_back(cnt); // 统计连续字符长度
                cnt = 0;
            }
        }

        int ans = 0;
        for (auto &a: groups) {
            if (a.empty()) continue;
            sort(a.rbegin(), a.rend());
            a.push_back(0);
            a.push_back(0); // 假设还有两个空串
            ans = max({ans, a[0] - 2, min(a[0] - 1, a[1]), a[2]});
        }
        return ans ? ans : -1;
    }
};
```

```go [sol-Go]
func maximumLength(s string) int {
	groups := [26][]int{}
	cnt := 0
	for i := range s {
		cnt++
		if i == len(s)-1 || s[i] != s[i+1] {
			groups[s[i]-'a'] = append(groups[s[i]-'a'], cnt) // 统计连续字符长度
			cnt = 0
		}
	}

	ans := 0
	for _, a := range groups {
		if len(a) == 0 {
			continue
		}
		slices.SortFunc(a, func(a, b int) int { return b - a })
		a = append(a, 0, 0) // 假设还有两个空串
		ans = max(ans, a[0]-2, min(a[0]-1, a[1]), a[2])
	}
	if ans == 0 {
		return -1
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $s$ 的长度。如果改用堆维护前三大可以做到 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。
