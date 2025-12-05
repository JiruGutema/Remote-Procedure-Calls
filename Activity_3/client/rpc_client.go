package main

import (
    "fmt"
    "log"
    "net/rpc"
    "time"
)

type Args struct {
    A, B int
}

func main() {
    client, err := rpc.Dial("tcp", "localhost:1234")
    if err != nil {
        log.Fatal("Error connecting to RPC server:", err)
    }

    args := Args{A: 10, B: 0} // Division by zero to trigger an error
    var reply int

    // Asynchronous RPC call
    call := client.Go("Calculator.Divide", &args, &reply, nil)

    select {
    case <-call.Done:
        if call.Error != nil {
            log.Println("RPC error:", call.Error)
        } else {
            fmt.Println("Result:", reply)
        }

    case <-time.After(2 * time.Second):
        log.Println("RPC call timed out")
    }
}
