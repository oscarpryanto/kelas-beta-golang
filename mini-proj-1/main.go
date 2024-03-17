package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Buku struct {
	kode          string
	judul         string
	pengarang     string
	penerbit      string
	jumlahHalaman int
	tahunTerbit   int
}

var listBuku []Buku

func InputMultiKata() string {
	var inputanMultiKata = bufio.NewReader(os.Stdin)
	input, error := inputanMultiKata.ReadString('\n')
	if error != nil {
		fmt.Print("Terjadi kesalahan pada saat input data buku : ", error)
		os.Exit(0)
	}
	input = strings.TrimSpace(input)
	return input
}

func TambahBuku() {

	inputJumlahHalaman := 0
	inputTahunTerbit := 0

	fmt.Println("===========================================")
	fmt.Println("Tambah Data Buku")
	fmt.Println("===========================================")

	// input kode
	fmt.Print("Silahkan Masukan Kode Buku : ")
	inputKodeBuku := InputMultiKata()

	// input judul
	fmt.Print("Silahkan Masukan Judul Buku : ")
	inputanJudulBuku := InputMultiKata()

	// input pengarang
	fmt.Print("Silahkan Masukan Pengarang Buku : ")
	inputanPengarangBuku := InputMultiKata()

	// input penerbit
	fmt.Print("Silahkan Masukan Penerbit Buku : ")
	inputanPenerbitBuku := InputMultiKata()

	// input jumlah halaman
	fmt.Print("Silahkan Masukan Jumlah Halaman Buku : ")
	_, error := fmt.Scanln(&inputJumlahHalaman)
	if error != nil {
		fmt.Print("Terjadi kesalahan pada saat menambahkan jumlah halaman buku : ", error)
		return
	}

	// input tahun terbit
	fmt.Print("Silahkan Masukan Tahun Terbit Buku : ")
	_, error = fmt.Scanln(&inputTahunTerbit)
	if error != nil {
		fmt.Print("Terjadi kesalahan pada saat menambahkan tahun terbit buku : ", error)
		return
	}

	listBuku = append(listBuku, Buku{
		kode:          inputKodeBuku,
		judul:         inputanJudulBuku,
		pengarang:     inputanPengarangBuku,
		penerbit:      inputanPenerbitBuku,
		jumlahHalaman: inputJumlahHalaman,
		tahunTerbit:   inputTahunTerbit,
	})
	fmt.Println("Berhasil Menambahkan data buku, kode: ", inputKodeBuku)

}

func TampilkanDaftarBuku() {
	CekEmptyList()
	fmt.Println("===========================================")
	fmt.Println("Data Daftar Buku")
	fmt.Println("===========================================")
	fmt.Println("")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	for urutan, item := range listBuku {
		fmt.Printf("No.%d | Kode buku: %s | Judul Buku: %s | Pengarang: %s | Penerbit: %s | Jumlah Halaman: %d | Tahun Terbit: %d\n",
			urutan+1,
			item.kode,
			item.judul,
			item.pengarang,
			item.penerbit,
			item.jumlahHalaman,
			item.tahunTerbit,
		)
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	}

}

func UbahBuku() {
	CekEmptyList()
	var indexBuku int
	var pilihanMenu int
	inputJumlahHalaman := 0
	inputTahunTerbit := 0

	fmt.Println("=================================")
	fmt.Println("Ubah Data Buku")
	fmt.Println("=================================")
	TampilkanDaftarBuku()
	fmt.Println("=================================")

	fmt.Print("Masukan kode Buku yang akan diubah : ")
	inputanUbahKodeBuku := InputMultiKata()
	indexBuku = FindIndexBook(listBuku, inputanUbahKodeBuku)

	if indexBuku < 0 {
		fmt.Println("Kode Buku Tidak Sesuai / tidak ditemukan!")
		UbahBuku()
		return
	}

	dataBukuSebelumEdit := listBuku[indexBuku]

	fmt.Println("=================================")
	fmt.Println("Silahkan Pilih elemen data pada buku yang ingin diubah : ")
	fmt.Println("1. Judul Buku")
	fmt.Println("2. Pengarang")
	fmt.Println("3. PenerBit")
	fmt.Println("4. Jumlah Halaman")
	fmt.Println("5. Tahun Terbit")
	fmt.Println("6. keluar")
	fmt.Println("===========================================")
	fmt.Print("Masukan Pilihan : ")

	_, err := fmt.Scanln(&pilihanMenu)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	switch pilihanMenu {
	case 1:
		fmt.Print("Silahkan Masukan Judul Buku : ")
		inputanJudulBuku := InputMultiKata()
		dataBukuSebelumEdit.judul = inputanJudulBuku
		fmt.Println("Judul Buku berhasil diubah ")
	case 2:
		fmt.Print("Silahkan Masukan Pengarang Buku : ")
		inputanPengarangBuku := InputMultiKata()
		dataBukuSebelumEdit.pengarang = inputanPengarangBuku
		fmt.Println("Pengarang Buku berhasil diubah ")
	case 3:
		fmt.Print("Silahkan Masukan Penerbit Buku : ")
		inputanPenerbitBuku := InputMultiKata()
		dataBukuSebelumEdit.penerbit = inputanPenerbitBuku
		fmt.Println("Penerbit Buku berhasil diubah ")
	case 4:
		fmt.Print("Silahkan Masukan Jumlah Halaman Buku : ")
		_, error := fmt.Scanln(&inputJumlahHalaman)
		if error != nil {
			fmt.Print("Terjadi kesalahan pada saat menambahkan jumlah halaman buku : ", error)
			return
		}
		dataBukuSebelumEdit.jumlahHalaman = inputJumlahHalaman
		fmt.Println("Jumlah halam buku berhasil diubah ")
	case 5:
		fmt.Print("Silahkan Masukan Tahun terbit Buku : ")
		_, error := fmt.Scanln(&inputTahunTerbit)
		if error != nil {
			fmt.Print("Terjadi kesalahan pada saat menambahkan Tahun terbit buku : ", error)
			return
		}
		dataBukuSebelumEdit.tahunTerbit = inputTahunTerbit
		fmt.Println("Jumlah tahun terbit buku berhasil diubah ")
	case 6:
		main()
	}

	listBuku[indexBuku] = dataBukuSebelumEdit
	TampilkanDaftarBuku()

}

func HapusBuku() {
	CekEmptyList()
	fmt.Println("=================================")
	fmt.Println("Hapus Data Buku")
	fmt.Println("=================================")
	TampilkanDaftarBuku()
	fmt.Println("=================================")
	var indexBuku int

	fmt.Print("Masukan Nomor Urut Buku yang akan dihapus : ")
	_, err := fmt.Scanln(&indexBuku)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	if (indexBuku-1) < 0 ||
		(indexBuku-1) > len(listBuku) {
		fmt.Println("Nomor urutan Buku Tidak Sesuai!")
		HapusBuku()
		return
	}

	dataKodeBuku := listBuku[indexBuku-1].kode
	listBuku = append(
		listBuku[:indexBuku-1],
		listBuku[indexBuku:]...,
	)

	fmt.Printf("Data buku dengan kode: %s Berhasil Dihapus!\n", dataKodeBuku)
}

func main() {
	// pilihanMenu := 0
	var pilihanMenu int
	fmt.Println("===========================================")
	fmt.Println("Aplikasi Manajemen Daftar Buku Perpustakaan")
	fmt.Println("===========================================")
	fmt.Println("Silahkan Pilih : ")
	fmt.Println("1. Tambah data Buku")
	fmt.Println("2. Lihat daftar buku")
	fmt.Println("3. Ubah data buku")
	fmt.Println("4. Hapus data buku")
	fmt.Println("5. Keluar")
	fmt.Println("===========================================")
	fmt.Print("Masukan Pilihan : ")

	_, err := fmt.Scanln(&pilihanMenu)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	switch pilihanMenu {
	case 1:
		TambahBuku()
	case 2:
		TampilkanDaftarBuku()
	case 3:
		UbahBuku()
	case 4:
		HapusBuku()
	case 5:
		os.Exit(0)
	}

	main()
}

func CekEmptyList() {
	if len(listBuku) <= 0 {
		fmt.Println("Data Buku Masih Kosong!, silahkan tambahkan data terlebih dahulu untuk mengakses menu ini")
		main()
	}
}

func FindIndexBook(listBuku []Buku, kodeBuku string) int {
	for i, buku := range listBuku {
		if buku.kode == kodeBuku {
			return i
		}
	}
	return -1
}
