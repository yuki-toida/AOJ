package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Watch
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())
	h := num / 3600
	m := (num % 3600) / 60
	s := (num % 3600) % 60
	fmt.Printf("%v:%v:%v\n", h, m, s)
}
