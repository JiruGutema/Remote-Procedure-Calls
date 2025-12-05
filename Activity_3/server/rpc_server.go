package main

import (
    "errors"
    "fmt"
    "net"
    "net/rpc"
)

type Args struct {
    A, B int
}

type Calculator struct{}

func (c *Calculator) Add(args *Args, reply *int) error {
    *reply = args.A + args.B
    return nil
}

func (c *Calculator) Subtract(args *Args, reply *int) error {
    *reply = args.A - args.B
    return nil
}

func (c *Calculator) Multiply(args *Args, reply *int) error {
    *reply = args.A * args.B
    return nil
}

func (c *Calculator) Divide(args *Args, reply *int) error {
    if args.B == 0 {
        return errors.New("cannot divide by zero")
    }
    *reply = args.A / args.B
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
