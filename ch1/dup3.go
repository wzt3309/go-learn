// Dup3 prints count and text of lines that appear once
// more in files, which by reading the entire file first
package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    count := make(map[string]int)
    filenames := os.Args[1:]
    for _, filename := range filenames {
        data, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
            continue
        }
        for _, line := range strings.Split(string(data), "\n") {
            count[line]++
        }
    }
    for line, n := range count {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
