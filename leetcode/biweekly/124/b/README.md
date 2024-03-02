[视频讲解](https://www.bilibili.com/video/BV1Sm411U7cR/)

## 提示 1

最后一次操作时，剩下的字母互不相同。

因为如果有相同字母，那么操作后还有剩余字母。

## 提示 2

设字母的最大出现次数为 $\textit{mx}$。

由于删除是从左到右进行的，最后剩下的就是出现次数等于 $\textit{mx}$ 的靠右字母（相同字母取出现位置最右的）。

```py [sol-Python3]
class Solution:
    def lastNonEmptyString(self, s: str) -> str:
        last = {c: i for i, c in enumerate(s)}
        cnt = Counter(s)
        mx = max(cnt.values())
        ids = sorted(last[ch] for ch, c in cnt.items() if c == mx)
        return ''.join(s[i] for i in ids)
```

```java [sol-Java]
class Solution {
    public String lastNonEmptyString(String S) {
        int[] cnt = new int[26];
        int[] last = new int[26];
        char[] s = S.toCharArray();
        for (int i = 0; i < s.length; i++) {
            int b = s[i] - 'a';
            cnt[b]++;
            last[b] = i;
        }

        // 注：也可以再遍历一次 s 直接得到答案，但效率不如下面，毕竟至多 26 个数
        List<Integer> ids = new ArrayList<>();
        int mx = Arrays.stream(cnt).max().orElseThrow();
        for (int i = 0; i < 26; i++) {
            if (cnt[i] == mx) {
                ids.add(last[i]);
            }
        }
        Collections.sort(ids);

        StringBuilder t = new StringBuilder(ids.size());
        for (int i : ids) {
            t.append(s[i]);
        }
        return t.toString();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string lastNonEmptyString(string s) {
        int cnt[26]{}, last[26]{};
        for (int i = 0; i < s.size(); i++) {
            int b = s[i] - 'a';
            cnt[b]++;
            last[b] = i;
        }

        // 注：也可以再遍历一次 s 直接得到答案，但效率不如下面，毕竟至多 26 个数
        vector<int> ids;
        int mx = ranges::max(cnt);
        for (int i = 0; i < 26; i++) {
            if (cnt[i] == mx) {
                ids.push_back(last[i]);
            }
        }
        ranges::sort(ids);

        string t(ids.size(), 0);
        for (int i = 0; i < ids.size(); i++) {
            t[i] = s[ids[i]];
        }
        return t;
    }
};
```

```go [sol-Go]
func lastNonEmptyString(s string) string {
	var cnt, last [26]int
	for i, b := range s {
		b -= 'a'
		cnt[b]++
		last[b] = i
	}

	// 注：也可以再遍历一次 s 直接得到答案，但效率不如下面，毕竟至多 26 个数
	ids := []int{}
	mx := slices.Max(cnt[:])
	for i, c := range cnt {
		if c == mx {
			ids = append(ids, last[i])
		}
	}
	slices.Sort(ids)

	t := make([]byte, len(ids))
	for i, id := range ids {
		t[i] = s[id]
	}
	return string(t)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|\log |\Sigma|)$，其中 $n$ 为 $s$ 的长度，$|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
