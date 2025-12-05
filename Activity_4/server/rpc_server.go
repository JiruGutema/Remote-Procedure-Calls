package main

import (
    "errors"
    "fmt"
    "net"
    "net/rpc"
    "sync"
)

type Args struct {
    A, B int
}

type Calculator struct {
    lastResult int
    mu         sync.Mutex // to safely handle concurrent clients
}

// Multiply
func (c *Calculator) Multiply(args *Args, reply *int) error {
    if args.A == 0 || args.B == 0 {
        return errors.New("multiplication by zero is not allowed")
    }
    result := args.A * args.B
    c.mu.Lock()
    c.lastResult = result
    c.mu.Unlock()
    *reply = result
    return nil
}

// Add
func (c *Calculator) Add(args *Args, reply *int) error {
    result := args.A + args.B
    c.mu.Lock()
    c.lastResult = result
    c.mu.Unlock()
    *reply = result
    return nil
}

// Subtract
func (c *Calculator) Subtract(args *Args, reply *int) error {
    result := args.A - args.B
    c.mu.Lock()
    c.lastResult = result
    c.mu.Unlock()
    *reply = result
    return nil
}

// Divide
func (c *Calculator) Divide(args *Args, reply *int) error {
    if args.B == 0 {
        return errors.New("cannot divide by zero")
    }
    result := args.A / args.B
    c.mu.Lock()
    c.lastResult = result
    c.mu.Unlock()
    *reply = result
    return nil
}

// New method: GetLastResult
func (c *Calculator) GetLastResult(args *Args, reply *int) error {
    c.mu.Lock()
    defer c.mu.Unlock()
    *reply = c.lastResult
    return nil
}

func main() {
    calc := new(Calculator)
    rpc.Register(calc)

    listener, err := net.Listen("tcp", ":1234")
    if err != nil {
        fmt.Println("Error starting RPC server:", err)
        return
    }

    fmt.Println("RPC server listening on port 1234...")
    rpc.Accept(listener)
}
