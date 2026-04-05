package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	// Variabel i yang bertipe integer berfungsi sebagai counter pada looping.
	// Variabel n yang bertipe integer berfungsi untuk menerima inputan angka dari pengguna.
	var i, n int

	// Array angka[] int berfungsi sebagai variabel yang akan digunakan pada RPC method sebagai argumen berjenis in/out.
	var angka []int

	// Memanggil RPC melalui protokol HTTP berdasarkan transport layer dan port yang telah ditentukan pada server.
	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	// Error handling pada RPC.
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	// Melakukan input angka pada variabel n.
	fmt.Print("Silahkan masukkan angka untuk batas: ")
	fmt.Scanf("%d", &n)

	// Memanggil method RpcPrime melalui tipe alias yang telah diregistrasi pada RPC beserta argumen-argumen yang diperlukan.
	client.Call("API.RpcPrime", n, &angka)

	// Looping untuk mengoutputkan response yang diberikan oleh RPC melalui server.
	for i = 0; i < len(angka); i++ {
		fmt.Print(angka[i], " ")
	}
}
