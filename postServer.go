package main
import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strconv"
    "time"
    "crypto/md5"
    "io"
)


func postApp(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    //注意:如果没有调用ParseForm方法, 下面无法获取表单数据

    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h, strconv.FormatInt(crutime, 10))
        token := fmt.Sprintf("%x", h.Sum(nil))

        fmt.Println(token)

        t, _ := template.ParseFiles("post.html")
        t.Execute(w, token)
    } else {

        token := r.Form.Get("token")
        if token != "abc" {
            fmt.Printf("token: %s\n", token)
        } else {
            fmt.Println("token not existed!")
        }


        fmt.Println("appname:", r.Form["appname"])
        fmt.Println("namespace:", r.Form["division"])
    }
}


func main() {
    http.HandleFunc("/post", postApp)

    err := http.ListenAndServe(":10000", nil)
    if err != nil {
        log.Fatal("ListenAdnServe:", err)
    }
}
