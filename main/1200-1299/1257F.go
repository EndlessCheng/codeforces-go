package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func Sol1257F(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	op := [30][100]int8{}
	baseOnes := [100]int8{}
	for j := 0; j < n; j++ {
		var v uint
		Fscan(in, &v)
		for i := range op {
			if v>>uint(i)&1 == 0 {
				op[i][j] = 1
			} else {
				op[i][j] = -1
			}
		}
		baseOnes[j] = int8(bits.OnesCount(v))
	}

	for ones := int8(0); ones <= 30; ones++ {
		mp := map[[100]int8]uint{}
		for i := uint(0); i < 1<<15; i++ {
			sum := [100]int8{}
			for j := uint(0); j < 15; j++ {
				if i>>j&1 == 1 {
					for k, v := range op[j] {
						sum[k] += v
					}
				}
			}
			mp[sum] = i
		}
		curOnes := [100]int8{}
		for i := 0; i < n; i++ {
			curOnes[i] = ones - baseOnes[i]
		}
		for i := uint(0); i < 1<<15; i++ {
			curOnesCopy := [100]int8{}
			for j, v := range curOnes {
				curOnesCopy[j] = v
			}
			for j := uint(0); j < 15; j++ {
				if i>>j&1 == 1 {
					for k, v := range op[j+15] {
						curOnesCopy[k] -= v
					}
				}
			}
			if v, ok := mp[curOnesCopy]; ok {
				Fprint(out, i<<15|v)
				return
			}
		}
	}
	Fprint(out, -1)
}

//func main() {
//	Sol1257F(os.Stdin, os.Stdout)
//}
