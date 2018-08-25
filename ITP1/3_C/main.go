package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Swapping two number
func main() {
	sc := bufio.NewScanner(os.Stdin)
	s := ""
	for i := 0; i < 3000; i++ {
		sc.Scan()
		ary := strings.Split(strings.TrimSpace(sc.Text()), " ")
		x, _ := strconv.Atoi(ary[0])
		y, _ := strconv.Atoi(ary[1])
		if x == 0 && y == 0 {
			break
		} else if x < y {
			s += strconv.Itoa(x) + " " + strconv.Itoa(y) + "\n"
		} else {
			s += strconv.Itoa(y) + " " + strconv.Itoa(x) + "\n"
		}
	}
	fmt.Print(s)
}
