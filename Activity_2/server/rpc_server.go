package main

import (
    "errors"
    "fmt"
    "net"
    "net/rpc"
)

// Args holds the arguments for arithmetic operations
type Args struct {
    A, B int
}

// Calculator provides methods for arithmetic operations
type Calculator int

// Multiply multiplies two integers
func (c *Calculator) Multiply(args *Args, reply *int) error {
    if args.A == 0 || args.B == 0 {
        return errors.New("multiplication by zero is not allowed")
    }
    *reply = args.A * args.B
    return nil
}

// Add adds two integers
func (c *Calculator) Add(args *Args, reply *int) error {
    *reply = args.A + args.B
    return nil
}

// Subtract subtracts B from A
func (c *Calculator) Subtract(args *Args, reply *int) error {
    *reply = args.A - args.B
    return nil
}

// Divide divides A by B
func (c *Calculator) Divide(args *Args, reply *int) error {
    if args.B == 0 {
        return errors.New("cannot divide by zero")
    }
    *reply = args.A / args.B
    return nil
}

func main() {
    // Register the Calculator service
    calc := new(Calculator)
    rpc.Register(calc)

    // Listen on TCP port 1234
    listener, err := net.Listen("tcp", ":1234")
    if err != nil {
        fmt.Println("Error starting RPC server:", err)
        return
    }
    fmt.Println("RPC server is listening on port 1234...")

    // rpc.Accept handles each client in a separate goroutine internally
    rpc.Accept(listener)
}
