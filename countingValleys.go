package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
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
        fmt.Println(lvl)
    }
    fmt.Print(lvl)

    return count
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    s := readLine(reader)

    result := countingValleys(n, s)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
