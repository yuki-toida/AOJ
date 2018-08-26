package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// シャッフル
func main() {
	sc := bufio.NewScanner(os.Stdin)
	s := ""
	for {
		sc.Scan()
		w := sc.Text()
		wc := len([]rune(w))
		if w == "-" {
			break
		}
		sc.Scan()
		m, _ := strconv.Atoi(sc.Text())
		for i := 0; i < m; i++ {
			sc.Scan()
			h, _ := strconv.Atoi(sc.Text())
			w = w[h:wc] + w[0:h]
		}
		s += w + "\n"
	}
	fmt.Print(s)
}
