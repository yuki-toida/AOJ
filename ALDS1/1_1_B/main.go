package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 最大公約数判定
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	ary := strings.Split(strings.TrimSpace(sc.Text()), " ")
	x, _ := strconv.Atoi(ary[0])
	y, _ := strconv.Atoi(ary[1])
	for 0 < y {
		if x < y {
			x, y = y, x
		}
		r := x % y
		if r == 0 {
			break
		}
		x = y
		y = r
	}
	fmt.Println(y)
}
