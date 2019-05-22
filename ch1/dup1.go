// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    count := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        count[input.Text()]++
    }

    for line, n := range count {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
