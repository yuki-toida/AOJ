package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// X Cubic
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())
	fmt.Println(num * num * num)
}
