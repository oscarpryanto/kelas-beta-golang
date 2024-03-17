package controller

import (
	"encoding/json"
	"fmt"
	"miniproject1/manajemen_buku_perpustakaan/model"
	"os"
	"sort"
	"sync"
)

// main function to display all book 
func TampilkanDaftarBuku() {
	fmt.Println("=================================")
	fmt.Println("Lihat Daftar Buku")
	fmt.Println("=================================")
	fmt.Println("Memuat data ...")
	ListBuku := LihatDaftarBuku()

	// sort book base on judul A-Z
	sort.Slice(ListBuku, func(i, j int) bool {
		return ListBuku[i].Judul < ListBuku[j].Judul
	})

	fmt.Println("")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	for urutan, item := range ListBuku {
		fmt.Printf("No.%d | Kode buku: %s | Judul Buku: %s | Pengarang: %s | Penerbit: %s | Jumlah Halaman: %d | Tahun Terbit: %d\n",
			urutan+1,
			item.Kode,
			item.Judul,
			item.Pengarang,
			item.Penerbit,
			item.JumlahHalaman,
			item.TahunTerbit,
		)
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	}

}

// helper function to read each file and parse json data into book struct through chanel
func lihatBuku(ch <-chan string, chBooks chan model.Buku, wg *sync.WaitGroup) {
	var buku model.Buku
	for kodeBuku := range ch {
		dataJSON, err := os.ReadFile(fmt.Sprintf("books/%s", kodeBuku))
		if err != nil {
			fmt.Println("Terjadi error:", err)
		}

		err = json.Unmarshal(dataJSON, &buku)
		if err != nil {
			fmt.Println("Terjadi error:", err)
		}

		chBooks <- buku
	}
	wg.Done()
}


// function to read all json file and then retunr list of book  
func LihatDaftarBuku() []model.Buku {
	ListBuku := []model.Buku{}

	listJsonBuku, err := os.ReadDir("books")
	if err != nil {
		fmt.Println("Terjadi error: ", err)
	}

	wg := sync.WaitGroup{}

	ch := make(chan string)
	chBooks := make(chan model.Buku, len(listJsonBuku))

	jumlahPetugas := 2

	for i := 0; i < jumlahPetugas; i++ {
		wg.Add(1)
		go lihatBuku(ch, chBooks, &wg)
	}

	for _, fileBuku := range listJsonBuku {
		ch <- fileBuku.Name()
	}

	close(ch)

	wg.Wait()

	close(chBooks)

	for dataBuku := range chBooks {
		ListBuku = append(ListBuku, dataBuku)
	}

	return ListBuku
}
