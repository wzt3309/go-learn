// Dup2 prints count and text of lines that appear more
// than once in the stdin or a list of files
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    count := make(map[string]int)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, count)
    } else {
         for _, file := range files {
             f, err := os.Open(file)
             if err != nil {
                 fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                 continue
             }
             countLines(f, count)
             f.Close()
         }
    }

    for line, n := range count {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

func countLines(f *os.File, count map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        count[input.Text()]++
    }
}
