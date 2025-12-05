package main

import (
    "fmt"
    "log"
    "net/rpc"
)

type Args struct {
    A, B int
}

func main() {
    client, err := rpc.Dial("tcp", "localhost:1234")
    if err != nil {
        log.Fatal("Error connecting to RPC server:", err)
    }

    args := Args{A: 7, B: 3}
    var reply int

    // Call Add
    client.Call("Calculator.Add", &args, &reply)
    fmt.Println("Add:", reply)

    // Call Multiply
    client.Call("Calculator.Multiply", &args, &reply)
    fmt.Println("Multiply:", reply)

    // Call GetLastResult
    client.Call("Calculator.GetLastResult", &Args{}, &reply)
    fmt.Println("Last Result:", reply)
}
