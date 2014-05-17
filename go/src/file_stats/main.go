package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
    "strings"
)

// Solution for Problem of the Day for 2014-05-13, File Stats
// http://www.problemotd.com/problem/file-stats/


func exit_on_error(err error) {
    if err != nil {
        fmt.Println("Error ", err)
        os.Exit(1)
    }
}


func get_file_stats(path string) map[string]string {
    file, err := os.Open(path)
    defer file.Close()
    exit_on_error(err)
    stats := make(map[string]string)
    stats["filename"] = file.Name()
    file_info, _ := file.Stat()
    stats["size"] = strconv.FormatInt(file_info.Size(), 10)
    stats["modified"] = file_info.ModTime().String()
    contents, err := ioutil.ReadAll(file)
    exit_on_error(err)
    str_contents := string(contents)
    stats["lines"] = strconv.Itoa(strings.Count(str_contents, "\n"))
    stats["words"] = strconv.Itoa(len(strings.Split(str_contents, " ")))
    return stats
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Error: Need a file path")
        os.Exit(1)
    }
    stats := get_file_stats(os.Args[1])
    fmt.Printf("\nStats for \"%s\"\n\n", stats["filename"])
    fmt.Printf("Words: %s, Lines: %s, Size: %s bytes,\nModified: %s\n",
               stats["words"], stats["lines"], stats["size"], stats["modified"])
}
