package main

import (
	"fmt"
)

func main() {
	fmt.Println(andrewSwitches(100, 100))
}

func andrewSwitches(n int, totalSwitches int) (res int32) {
	count := 0

	for i := 1; i < totalSwitches+1; i++ {
		count += totalSwitches / i
	}
	res = int32(count)
	return res
}
