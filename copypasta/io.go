package copypasta

import (
	"bufio"
	. "fmt"
	"os"
)

func simpleIO() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	Fscan(in, &n)
}
