package main
import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, playground")
	digitSeparator(1345679)
}

func digitSeparator(n int32) {
    s := strconv.Itoa(int(n))
    for k, _ := range s {
	     charLeft := string(s[:len(s)-k])
	     fmt.Println(charLeft)
    }
}
