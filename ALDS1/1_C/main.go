package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// 素数判定
func main() {
	sc := bufio.NewScanner(os.Stdin)
	n := stdin(sc)
	cnt := 0
	for i := 0; i < n; i++ {
		if isPrime(stdin(sc)) {
			cnt++
		}
	}
	fmt.Println(cnt)
}

func stdin(sc *bufio.Scanner) int {
	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())
	return num
}

func isPrime(v int) bool {
	if v < 2 {
		return false
	}
	if v == 2 {
		return true
	}
	if v%2 == 0 {
		return false
	}
	n := int(math.Sqrt(float64(v)))
	for i := 3; i <= n; i += 2 {
		if v%i == 0 {
			return false
		}
	}
	return true
}
