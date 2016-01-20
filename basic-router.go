package main
import (
    "fmt"
    "html"
    "net/http"
    "log"
    "github.com/gorilla/mux"
)

func main() {
//    router := mux.NewRouter().StrictSlash(true)
    router := mux.NewRouter().StrictSlash(false)
    router.HandleFunc("/", Index)
    log.Fatal(http.ListenAndServe(":8080", router))
 
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.Path))
}
