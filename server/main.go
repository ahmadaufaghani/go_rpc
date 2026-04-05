package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// Tipe alias untuk mengubah fungsi menjadi method yang dapat diekspor dan digunakan melalui RPC.
type API int

// Method RpcPrime yang akan diekspor sehingga dapat dipanggil melalui RPC.
func (a *API) RpcPrime(n int, reply *[]int) error {
	// Variabel i dan j yang bertipe integer berfungsi sebagai counter pada looping, serta total untuk menghitung jumlah bilangan yang dapat membagi angka yang sedang dilakukan pengecekan apakah bilangan tersebut prima atau bukan.
	var i, j, total int

	// Array angka bertipe integer berfungsi untuk menyimpan angka mana saja yang termasuk bilangan prima berdasarkan batas angka maksimal yang diberikan oleh pengguna.
	var angka []int

	// Looping angka hingga batas angka maksimal yang diberikan oleh pengguna.
	for i = 0; i < n; i++ {
		total = 0

		// Looping untuk menghitung berapa banyak jumlah bilangan yang dapat membagi habis angka yang sekarang sedang dilakukan pengecekan (mulai dari angka-angka sebelumnya hingga dibagi oleh angka itu sendiri).
		for j = 1; j <= i+1; j++ {
			if (i+1)%j == 0 {
				total = total + 1
			}
		}

		// Melakukan pengecekan pada angka yang sedang diperiksa, jika jumlah bilangan yang dapat membagi habis pada variabel total sama dengan 2, maka angka tersebut akan ditambahkan pada Array var angka []int, karena termasuk bilangan prima.
		if total == 2 {
			angka = append(angka, i+1)
		}
	}

	// Array var angka[]int yang berisi kumpulan bilangan prima sesuai batas angka yang diberikan oleh pengguna, akan dikembalikan melalui atribut method reply *[]int yang berjenis in/out.
	*reply = angka
	return nil
}

func main() {
	// Mengalokasi memori untuk tipe alias dan pointer yang baru dialokasi untuk tipe data tersebut disimpan pada variabel api.
	var api = new(API)

	// RPC melakukan registrasi terhadap variabel tersebut yang berisi alokasi memori dan pointer tipe data alias, sehingga apabila ingin melakukan pemanggilan method dapat dilakukan melalui tipe alias tersebut.
	err := rpc.Register(api)

	// Error handling pada RPC.
	if err != nil {
		log.Fatal("error registering API", err)
	}

	// Mengatur RPC agar dapat melakukan request dan response pada client melalui protokol HTTP pada server yang sedang dijalankan.
	rpc.HandleHTTP()

	// Mengatur server agar berjalan pada port 4040 dan menggunakan transport layer TCP yang kemudian disimpan pada variabel listener.
	listener, err := net.Listen("tcp", ":4040")

	// Log untuk menginformasikan bahwa server untuk RPC telah berjalan pada port 4040.
	log.Printf("running rpc on port %d", 4040)

	// Menjalankan server yang telah disiapkan berdasarkan pengaturan sebelumnya pada variabel listener menggunakan modul http.
	err = http.Serve(listener, nil)

	// Error handling pada http.
	if err != nil {
		log.Fatal("error running: ", err)
	}

}
