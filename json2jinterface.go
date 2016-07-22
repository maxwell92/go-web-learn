package main

import (
    "fmt"
    "encoding/json"
)

/* bool    ----> json booleans
 * float64 ----> josn numbers
 * string  ----> json strings
 * nil     ----> json null
 */


b := []byte(`{"Name":"Wednesday", "Age":6, "Parent":["Gomez", "Morticia"]}`)

var f interface{}
err := json.Unmarshal(b, &f)

/*
 * f = map[string]interface{}{
 *     "Name": "Wednesday",
 *     "Age": 6,
 *     "Parents": []interfaces{}{
 *         "Gomez",
 *         "Morticia",
 *      }
 * }
 */

m := f.(map[string]interface{})

for k, v := range m {
    switch vv := v.(type) {
        case string:
            fmt.Println(k, "is string", vv)
        case int:
            fmt.Println(k, "is int", vv)
        case float64:
            fmt.Println(k, "is float64", vv)
        case []interface{}:
            fmt.Println(k, "is an array:")
            for i, u := range vv {
            fmt.Println(i, u)
            }
        default:
            fmt.Println(k, "is of a type I don't know how to handle")

    }
}

/* -------- upper is the official handling strategy ---------- */

js, err := NewJson([]byte(`{
    "test": {
        "array": [1, "2", 3],
        "int": 10,
        "float": 5.150,
        "bignum": 111111123123123123123123,
        "string": "simplejson",
        "bool": true

    }

}`))

arr, _ := js.Get("test").Get("array").Array()
i, _ := js.Get("test").Get("int").Int()
ms := js.Get("test").Get("string").MushString()

/* -------- upper is the third-party handling strategy ---------- */
/* -------- https://github.com/bitly/go-simplejson ---------- */



