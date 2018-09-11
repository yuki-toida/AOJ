package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 組み合わせの数
func main() {
	sc := bufio.NewScanner(os.Stdin)
	for {
		sc.Scan()
		ary := strings.Split(sc.Text(), " ")
		n, _ := strconv.Atoi(ary[0])
		x, _ := strconv.Atoi(ary[1])
		if n == 0 && x == 0 {
			break
		}
		cnt := 0
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if j <= i {
					continue
				}
				for k := 1; k <= n; k++ {
					if k <= i || k <= j {
						continue
					} else if i+j+k == x {
						cnt++
					}
				}
			}
		}
		fmt.Println(cnt)
	}
}
