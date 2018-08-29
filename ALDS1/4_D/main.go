package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var n, k int
var w []int

// 割り当て
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	str := strings.Split(sc.Text(), " ")
	n, _ = strconv.Atoi(str[0])
	k, _ = strconv.Atoi(str[1])
	w = make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		w[i], _ = strconv.Atoi(sc.Text())
	}
	p := search()
	fmt.Println(p)
}

func max(p int) int {
	i := 0
	for j := 0; j < k; j++ {
		sum := 0
		for sum+w[i] <= p {
			sum += w[i]
			i++
			if i == n {
				return n
			}
		}
	}
	return i
}

func search() int {
	var left, mid int
	right := 10000 * 100000
	for 1 < right-left {
		mid = (right + left) / 2
		p := max(mid)
		if n <= p {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}
