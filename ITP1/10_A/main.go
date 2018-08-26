package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// 距離
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	ary := strings.Split(sc.Text(), " ")
	x1, _ := strconv.ParseFloat(ary[0], 64)
	y1, _ := strconv.ParseFloat(ary[1], 64)
	x2, _ := strconv.ParseFloat(ary[2], 64)
	y2, _ := strconv.ParseFloat(ary[3], 64)
	fmt.Printf("%v\n", math.Sqrt(math.Pow((x2-x1), 2)+math.Pow((y2-y1), 2)))
}
