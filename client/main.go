package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	var i, n int
	var angka []int
	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	fmt.Print("Silahkan masukkan angka untuk batas: ")
	fmt.Scanf("%d", &n)

	client.Call("API.RpcPrime", n, &angka)

	for i = 0; i < len(angka); i++ {
		fmt.Print(angka[i], " ")
	}
}
