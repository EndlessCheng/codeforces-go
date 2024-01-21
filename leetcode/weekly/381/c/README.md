[视频讲解](https://www.bilibili.com/video/BV1Q5411C7mN/) 第三题。

统计每个字母的出现次数，按照出现次数从大到小排序。

根据 [排序不等式](https://baike.baidu.com/item/%E6%8E%92%E5%BA%8F%E4%B8%8D%E7%AD%89%E5%BC%8F/7775728)，出现次数前 $8$ 大的字母，只需要按一次；出现次数前 $9$ 到 $16$ 大的字母，需要按两次；依此类推。

把出现次数和对应的按键次数相乘再相加，得到的按键次数之和就是最小的。

```py [sol-Python3]
class Solution:
    def minimumPushes(self, word: str) -> int:
        a = sorted(Counter(word).values(), reverse=True)
        return sum(c * (i // 8 + 1) for i, c in enumerate(a))
```

```java [sol-Java]
class Solution {
    public int minimumPushes(String word) {
        int[] cnt = new int[26];
        for (char b : word.toCharArray()) {
            cnt[b - 'a']++;
        }
        Arrays.sort(cnt);

        int ans = 0;
        for (int i = 0; i < 26; i++) {
            ans += cnt[25 - i] * (i / 8 + 1);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumPushes(string word) {
        int cnt[26]{};
        for (char b: word) {
            cnt[b - 'a']++;
        }
        ranges::sort(cnt, greater<int>());

        int ans = 0;
        for (int i = 0; i < 26; i++) {
            ans += cnt[i] * (i / 8 + 1);
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumPushes(word string) (ans int) {
	cnt := [26]int{}
	for _, b := range word {
		cnt[b-'a']++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(cnt[:])))

	for i, c := range cnt {
		ans += c * (i/8 + 1)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+|\Sigma|\log |\Sigma|)$，其中 $n$ 为 $\textit{word}$ 的长度，$|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$O(|\Sigma|)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
