// Run it in go playground
package main

import (
    "fmt"
)

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
    var count int32 = 0
    var lvl int32 = 0
    for i:=0; i < len(s); i++ {

        if string(s[i]) == "D" {
            lvl -= 1
        } else if string(s[i]) == "U" {
            lvl += 1
        } 
        if lvl == 0 && string(s[i]) == "U" {
            count += 1
        }
    }

    return count
}

func main() {
    n := int32(8)
    s := "DDUUDDD"

    result := countingValleys(n, s)

    fmt.Print(result)
}
