package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Print test case
func main() {
	sc := bufio.NewScanner(os.Stdin)
	s := ""
	for i := 1; i <= 10000; i++ {
		sc.Scan()
		num, _ := strconv.Atoi(sc.Text())
		if num == 0 {
			break
		}
		s += fmt.Sprintf("Case %v: %v\n", i, num)
	}
	fmt.Print(s)
}
