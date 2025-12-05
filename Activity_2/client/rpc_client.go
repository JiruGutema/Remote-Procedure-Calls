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
    // Connect to RPC server
    client, err := rpc.Dial("tcp", "localhost:1234")
    if err != nil {
        log.Fatal("Error connecting to RPC server:", err)
    }

    args := Args{A: 20, B: 5}
    var reply int

    // Call Add
    err = client.Call("Calculator.Add", &args, &reply)
    if err != nil {
        log.Fatal("Add Error:", err)
    }
    fmt.Println("Add:", reply)

    // Call Subtract
    err = client.Call("Calculator.Subtract", &args, &reply)
    if err != nil {
        log.Fatal("Subtract Error:", err)
    }
    fmt.Println("Subtract:", reply)

    // Call Multiply
    err = client.Call("Calculator.Multiply", &args, &reply)
    if err != nil {
        log.Fatal("Multiply Error:", err)
    }
    fmt.Println("Multiply:", reply)

    // Call Divide
    err = client.Call("Calculator.Divide", &args, &reply)
    if err != nil {
        log.Fatal("Divide Error:", err)
    }
    fmt.Println("Divide:", reply)
}
