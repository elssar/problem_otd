package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter text: ")
    text, _ := reader.ReadString('\n')
    frequency := make(map[uint8]int)
    for i := 0; i < len(text); i++ {
        if frequency[text[i]] != 0 {
            frequency[text[i]] += 1
        } else {
            frequency[text[i]] = 1
        }
    }
    for k, v := range(frequency) {
        if k != '\n' {
            count_in_dots := strings.Repeat(".", v)
            fmt.Printf("%q:\t%q\n", k, count_in_dots)
        }
    }
}
