package main
import (
    "fmt"
    "net/http"
)
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello SRE! This is deployed by CI/CD Pipeline!!!")
    })
    http.ListenAndServe(":8080", nil)
}
