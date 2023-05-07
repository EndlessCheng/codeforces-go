### 本题视频讲解

见[【周赛 344】](https://www.bilibili.com/video/BV1YL41187Rx/)第二题，欢迎点赞投币！

### 思路

用哈希表 $\textit{cnt}$ 统计每个数的出现次数。

用哈希表 $\textit{freq}$ 统计出现次数的出现次数，从而可以 $\mathcal{O}(1)$ 回答 $\texttt{hasFrequency}$。

添加删除元素的时候，除了修改 $\textit{cnt}[\textit{number}]$，还需要根据 $\textit{cnt}[\textit{number}]$ 的变化来修改 $\textit{freq}$，具体见代码。

```py [sol1-Python3]
class FrequencyTracker:
    def __init__(self):
        self.cnt = Counter()  # 每个数的出现次数
        self.freq = Counter()  # 出现次数的出现次数

    def add(self, number: int) -> None:
        self.freq[self.cnt[number]] -= 1  # 直接减，因为下面询问的不会涉及到 frequency=0
        self.cnt[number] += 1
        self.freq[self.cnt[number]] += 1

    def deleteOne(self, number: int) -> None:
        if self.cnt[number] == 0: return  # 不删除任何内容
        self.freq[self.cnt[number]] -= 1
        self.cnt[number] -= 1
        self.freq[self.cnt[number]] += 1

    def hasFrequency(self, frequency: int) -> bool:
        return self.freq[frequency] > 0
```

```java [sol1-Java]
class FrequencyTracker {
    private Map<Integer, Integer> cnt = new HashMap<>(); // 每个数的出现次数
    private Map<Integer, Integer> freq = new HashMap<>(); // 出现次数的出现次数

    public FrequencyTracker() {}

    public void add(int number) {
        // 直接减，因为下面询问的不会涉及到 frequency=0
        freq.merge(cnt.getOrDefault(number, 0), -1, Integer::sum);
        int c = cnt.merge(number, 1, Integer::sum);
        freq.merge(c, 1, Integer::sum);
    }

    public void deleteOne(int number) {
        if (cnt.getOrDefault(number, 0) == 0) return; // 不删除任何内容
        freq.merge(cnt.get(number), -1, Integer::sum);
        int c = cnt.merge(number, -1, Integer::sum);
        freq.merge(c, 1, Integer::sum);
    }

    public boolean hasFrequency(int frequency) {
        return freq.getOrDefault(frequency, 0) > 0;
    }
}
```

```cpp [sol1-C++]
class FrequencyTracker {
    unordered_map<int, int> cnt; // 每个数的出现次数
    unordered_map<int, int> freq; // 出现次数的出现次数
public:
    FrequencyTracker() {}

    void add(int number) {
        --freq[cnt[number]]; // 直接减，因为下面询问的不会涉及到 frequency=0
        ++freq[++cnt[number]];
    }

    void deleteOne(int number) {
        if (!cnt[number]) return; // 不删除任何内容
        --freq[cnt[number]];
        ++freq[--cnt[number]];
    }

    bool hasFrequency(int frequency) {
        return freq[frequency];
    }
};
```

```go [sol1-Go]
type FrequencyTracker struct {
	cnt  map[int]int // 每个数的出现次数
	freq map[int]int // 出现次数的出现次数
}

func Constructor() (_ FrequencyTracker) {
	return FrequencyTracker{map[int]int{}, map[int]int{}}
}

func (f FrequencyTracker) Add(number int) {
	f.freq[f.cnt[number]]-- // 直接减，因为下面询问的不会涉及到 frequency=0
	f.cnt[number]++
	f.freq[f.cnt[number]]++
}

func (f FrequencyTracker) DeleteOne(number int) {
	if f.cnt[number] == 0 {
		return // 不删除任何内容
	}
	f.freq[f.cnt[number]]--
	f.cnt[number]--
	f.freq[f.cnt[number]]++
}

func (f FrequencyTracker) HasFrequency(frequency int) bool {
	return f.freq[frequency] > 0
}
```

### 复杂度分析

- 时间复杂度：所有操作均为 $\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(q)$。其中 $q$ 为操作次数。
