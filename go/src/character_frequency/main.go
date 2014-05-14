package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strings"
)

// Solution for Problem of the Day for 2014-05-12, Character Frequency
// http://www.problemotd.com/problem/character-frequency/

type CharFreq struct {
    Char rune
    Count int
}

type ByCount []CharFreq

func (a ByCount) Len() int {
    return len(a)
}

func (a ByCount) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

// Reversed, so that sort is in descending order
func (a ByCount) Less(i, j int) bool {
    return a[i].Count > a[j].Count
}

func InCharFreq(this []CharFreq, x rune) int {
    i, j := 0, len(this)
    for i < j {
        if this[i].Char == x {
            return i
        }
        i += 1
    }
    return -1
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter text: ")
    text, _ := reader.ReadString('\n')
    frequency := []CharFreq{}

    for _, x := range text {
        index := InCharFreq(frequency, x)
        if index < 0 {
            frequency = append(frequency, CharFreq{x, 1})
        } else {
            frequency[index].Count += 1
        }
    }

    sort.Sort(ByCount(frequency))

    for _, v := range(frequency) {
        if v.Char != '\n' {
            count_in_dots := strings.Repeat(".", v.Count)
            fmt.Printf("%q:\t%q\n", v.Char, count_in_dots)
        }
    }
}
