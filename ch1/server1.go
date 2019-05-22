// Server1 is a minimal "echo" server
package main
import (
    "fmt"
    "net/http"
    "log"
)

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "URL.path=%q\n", r.URL.Path)
}
