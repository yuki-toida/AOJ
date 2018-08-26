package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 表計算
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	ary := strings.Split(sc.Text(), " ")
	r, _ := strconv.Atoi(ary[0])
	c, _ := strconv.Atoi(ary[1])

	hsum := make([]int, c+1)
	s := ""
	for i := 0; i < r; i++ {
		sc.Scan()
		ary := strings.Split(sc.Text(), " ")
		rsum := 0
		for j := 0; j < c; j++ {
			v, _ := strconv.Atoi(ary[j])
			rsum += v
			s += ary[j] + " "
			hsum[j] += v
		}
		s += fmt.Sprintf("%v\n", rsum)
		hsum[c] += rsum
	}
	for _, v := range hsum {
		s += fmt.Sprintf("%v ", v)
	}
	fmt.Println(strings.TrimSpace(s))
}
