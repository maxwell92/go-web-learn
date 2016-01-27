package main
import (
    "fmt"
    "html"
    "net/http"
    "log"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/mushroom", Mushroom)
    router.HandleFunc("/magic/{magicID}", MagicShow)
    
    log.Fatal(http.ListenAndServe(":8080", router))
 
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.Path))
}

func Mushroom(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello mushroomo")
}

func MagicShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    magicID := vars["magicID"]
    fmt.Fprintln(w, "MagicShow", magicID)
}
