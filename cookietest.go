package main

import (
    "fmt"
    "net/http"
    "time"
)

func cookie(w http.ResponseWriter, r *http.Request) {
    expiration := time.Now()
    expiration = expiration.AddDate(1, 0, 0)
    http.Cookie{Name: "username", Value: "mushroom", Expires: expiration}
    http.SetCookie(w, &cookie)
}

func readcookie(w http.ResponseWriter, r *http.Request) {
    cookie, _ := r.Cookie("username")
    fmt.Fprint(w, cookie)

    // or

    for _, cookie := range r.Cookies() {
        fmt.Fprint(w, cookie.Name)
    }
}

func main() {

}
