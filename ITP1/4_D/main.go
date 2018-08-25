package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Max min sum
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	max := -1000000
	min := 1000000
	sum := int64(0)
	for i := 0; i < n; i++ {
		sc.Scan()
		num, _ := strconv.Atoi(sc.Text())
		if max < num {
			max = num
		}
		if num < min {
			min = num
		}
		sum += int64(num)
	}
	fmt.Printf("%v %v %v\n", min, max, sum)
}
