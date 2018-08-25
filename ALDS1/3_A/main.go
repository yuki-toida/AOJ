package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ary []int

// スタック
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	for _, v := range strings.Split(strings.TrimSpace(sc.Text()), " ") {
		n, err := strconv.Atoi(v)
		if err == nil {
			ary = append(ary, n)
		} else {
			switch v {
			case "+":
				p1 := pop()
				p2 := pop()
				ary = append(ary, p2+p1)
			case "-":
				p1 := pop()
				p2 := pop()
				ary = append(ary, p2-p1)
			case "*":
				p1 := pop()
				p2 := pop()
				ary = append(ary, p2*p1)
			}
		}
	}
	fmt.Println(ary[0])
}

func pop() int {
	i := len(ary) - 1
	n := ary[i]
	ary = ary[:i]
	return n
}
