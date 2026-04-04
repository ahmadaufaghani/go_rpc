package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type API int

func (a *API) RpcPrime(n int, reply *[]int) error {
	var i, j, total int
	var angka []int
	for i = 0; i < n; i++ {
		total = 0
		for j = 1; j <= i+1; j++ {
			if (i+1)%j == 0 {
				total = total + 1
			}
		}
		if total == 2 {
			angka = append(angka, i+1)
		}
	}

	*reply = angka
	return nil
}

func main() {
	var api = new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	log.Printf("running rpc on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("error running: ", err)
	}

}
