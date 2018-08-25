package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 大小関係
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	ary := strings.Split(strings.TrimSpace(sc.Text()), " ")
	a, _ := strconv.Atoi(ary[0])
	b, _ := strconv.Atoi(ary[1])
	if a < b {
		fmt.Println("a < b")
	} else if a == b {
		fmt.Println("a == b")
	} else {
		fmt.Println("a > b")
	}
}
