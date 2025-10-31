package P2000_2999

import (
	. "fmt"
	"io"
	"math/big"
)

// https://space.bilibili.com/206214
func p2152(in io.Reader, out io.Writer) {
	a, b := new(big.Int), new(big.Int)
	Fscan(in, a, b)
	Fprint(out, a.GCD(nil, nil, a, b))
}

//func main() { p2152(bufio.NewReader(os.Stdin), os.Stdout) }
