package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// How many divider
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	ary := strings.Split(strings.TrimSpace(sc.Text()), " ")
	a, _ := strconv.Atoi(ary[0])
	b, _ := strconv.Atoi(ary[1])
	c, _ := strconv.Atoi(ary[2])
	cnt := 0
	for i := a; i <= b; i++ {
		if c%i == 0 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
