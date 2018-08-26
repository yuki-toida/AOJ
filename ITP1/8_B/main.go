package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 数字の和
func main() {
	sc := bufio.NewScanner(os.Stdin)
	s := ""
	for {
		sc.Scan()
		str := sc.Text()
		if str == "0" {
			break
		}
		sum := 0
		for _, v := range str {
			n, err := strconv.Atoi(string(v))
			if err != nil {
				continue
			}
			sum += n
		}
		s += fmt.Sprintf("%v\n", sum)
	}
	fmt.Print(s)
}
