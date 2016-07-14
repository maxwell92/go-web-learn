package main

import (
    "fmt"
    "net/rpc"
    "os"
)

func main() {
    serverAddress := os.Args[1]
    client, err := rpc.DialHTTP("tcp", serverAddress+":10000")
    if err != nil {
        fmt.Println(err)
    }

    var result int
    err = client.Call("Myop.PrintAndReturn", "mushroom", &result)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(result)


}
