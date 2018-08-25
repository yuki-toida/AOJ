package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 公舎の入居者数
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	var all = [4][3][]int{}
	for i, x := range all {
		for j := range x {
			all[i][j] = make([]int, 10)
		}
	}
	for i := 0; i < n; i++ {
		sc.Scan()
		ary := strings.Split(sc.Text(), " ")
		b, _ := strconv.Atoi(ary[0])
		f, _ := strconv.Atoi(ary[1])
		r, _ := strconv.Atoi(ary[2])
		v, _ := strconv.Atoi(ary[3])
		all[b-1][f-1][r-1] += v
	}
	s := ""
	for i, x := range all {
		for j, y := range x {
			for k, z := range y {
				if k == 0 {
					s += fmt.Sprintf(" %v ", z)
				} else if k == 9 {
					s += fmt.Sprintf("%v", z)
				} else {
					s += fmt.Sprintf("%v ", z)
				}
			}
			if i == 3 && j == 2 {
			} else {
				s += "\n"
			}
		}
		if i < 3 {
			s += "####################\n"
		}
	}
	fmt.Println(s)
}
