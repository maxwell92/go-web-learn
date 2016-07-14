package main
import (
    "fmt"
    "net/rpc"
    "net/http"
    "errors"
)

type Myop int

func (op *Myop) PrintAndReturn(args string, result *int) error {
    re := 1
    *result = re
    if args == "" {
        return errors.New("WRONG")
    }
    fmt.Println(args)
    fmt.Println("received request")

    return nil
}

func main() {

    mush := new(Myop)
    rpc.Register(mush)
    rpc.HandleHTTP()

    err := http.ListenAndServe(":10000", nil)
    if err != nil {
        fmt.Println(err)
    }



}
