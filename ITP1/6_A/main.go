package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 数列の反転
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	s := ""
	for i := 0; i < n; i++ {
		sc.Scan()
		s = sc.Text() + " " + s
	}
	fmt.Println(strings.TrimSpace(s))
}
