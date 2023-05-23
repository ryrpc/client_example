package main

import (
        "encoding/json"
        "fmt"

        rycli "github.com/ryrpc/client"
)

type sumRequest struct {
        A int `json:"a"`
        B int `json:"b"`
}

func main() {

        var data json.Number
        client := rycli.NewClient()

        client.SetBaseURL("http://127.0.0.1:8080")

        dst := sumRequest{}
        dst.A = 12
        dst.B = 6
        err := client.Call("sum_struct", dst, &data)
        if err != nil {
                panic(err)
        }

        //data = res.Result.(json.Number)
        fmt.Println("data = ", data)
}
