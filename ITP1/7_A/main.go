package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 成績
func main() {
	sc := bufio.NewScanner(os.Stdin)
	s := ""
	for i := 0; i < 50; i++ {
		sc.Scan()
		ary := strings.Split(sc.Text(), " ")
		m, _ := strconv.Atoi(ary[0])
		f, _ := strconv.Atoi(ary[1])
		r, _ := strconv.Atoi(ary[2])
		if m == -1 && f == -1 && r == -1 {
			break
		}
		if m == -1 || f == -1 {
			s += "F\n"
		} else if 80 <= m+f {
			s += "A\n"
		} else if 65 <= m+f && m+f < 80 {
			s += "B\n"
		} else if 50 <= m+f && m+f < 65 {
			s += "C\n"
		} else if m+f < 30 {
			s += "F\n"
		} else {
			if 50 <= r {
				s += "C\n"
			} else {
				s += "D\n"
			}
		}
	}
	fmt.Print(s)
}
