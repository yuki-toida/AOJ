package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Rectangle
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	ary := strings.Split(strings.TrimSpace(sc.Text()), " ")
	x, _ := strconv.Atoi(ary[0])
	y, _ := strconv.Atoi(ary[1])

	fmt.Printf("%v %v\n", x*y, (x+y)*2)
}
