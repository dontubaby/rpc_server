// RPC server
// DONE!
package main

import (
	"fmt"
	"math"
	"net"
	"net/http"
	"net/rpc"
)

type Point struct {
	X, Y float64
}

type Points struct {
	A, B Point
}

type RPCserver int64

func (s *RPCserver) Dist(p Points, resp *float64) error {
	*resp = math.Sqrt(math.Pow(p.A.X-p.B.X, 2.0) + math.Pow(p.A.Y-p.B.Y, 2.0))
	return nil
}

func main() {

	rpcserver := new(RPCserver)
	rpc.Register(rpcserver)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Start listen")
	http.Serve(listener, nil)

}
