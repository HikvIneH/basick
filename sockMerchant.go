package main

import (
    "fmt"
    "strconv"
    "strings"
)

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
    var res int32 = 0;
    var temp []int32 
    count := 1

    for k, v := range ar {
        exists := false
        for _, v1 := range temp {
            if v1 == v {
                exists = true
                break
            }
        }
        if exists {
            continue
        }
        for j := int32(k+1); j < n; j++ {
            if v == ar[j] {
                temp = append(temp, v)
                count += 1
            }
        }
        res += int32(count/2)
        count = 1
    }

    return res;
}

func main() {
    n := int32(9)

    arTemp := strings.Split("10 20 20 10 10 30 50 10 20", " ")

    var ar []int32

    for i := 0; i < int(n); i++ {
        arItemTemp, _ := strconv.ParseInt(arTemp[i], 10, 64)
        arItem := int32(arItemTemp)
        ar = append(ar, arItem)
    }

    result := sockMerchant(n, ar)

    fmt.Print(result)
}
