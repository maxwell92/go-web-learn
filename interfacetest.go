package main

import (
    "fmt"
)

func hello() {
    fmt.Println("Hello")
}

type myInterface interface {
    hello()
}

func main() {

}
