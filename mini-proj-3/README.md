# MINI PROJECT 3 -> Extended Features from Mini Project 2

### Nama aplikasi         : Aplikasi CLI Manajemen Daftar Buku Perpustakaan
Fitur pada apliaksi	:

1. Semua fitur yang sudah ada pada Mini project 2 dengan update:
    - Fitur tambah buku yang sebelumnya disimpan kedalam bentuk **JSON** file, sekarang disimpan kedalam **database**, tidak lagi disimpan dalam **JSON** file.
    - Fitur update yang sebelumnya memanipulasi data di **JSON** file, sekarang memanipulasi data di **database**
    - Fitur delete yang sebelumnya menghapus file **JSON**, sekarang menerapkan mekanisme ***soft-delete*** di **database.**
    - Fitur list buku mengambil data dari **database**.
    - Jika ada fitur lain seperti **Print Buku** yang masih membaca data dari **JSON** file, jangan lupa untuk disesuaikan agar membaca dari data di database.
    - Isi **Struct** dari **Book**, disesuaikan seperti berikut ni :
        
        ```go
        ID        uint           `gorm:"primarykey" json:"id"`
        CreatedAt time.Time      `json:"created_at"`
        UpdatedAt time.Time      `json:"updated_at"`
        DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
        ISBN      string         `json:"isbn"`
        Penulis   string         `json:"penulis"`
        Tahun     uint           `json:"tahun"`
        Judul     string         `json:"judul"`
        Gambar    string         `json:"gambar"`
        Stok      uint           `json:"stok"`
        
        ```
        
2. Menambahkan fitur **import data** dari file **CSV**, dimana user harus **menginputkan path/lokasi** file csv tersebut, lalu meload file CSV tersebut dan isi datanya dimasukan kedalam database
3. Mekanisme **import data** bersifat ***upsert***, artinya jika data sudah ada, maka dia hanya mengupdate data didalam tabel berdasarkan ID, jika tidak, dia akan memasukan data seperti biasa
4. Mengimplementasi ORM library **GORM**
5. Mengimplementasi ***auto migrate*** agar tabel terbuat otomatis ketika project dijalankan
6. Menerapkan unit testing untuk fitur **Tambah**, **Update**, **List**, dan **Hapus**
7. Unit test ketika dijalankan dengan ketentuan :
    - Database kosong, tidak ada tabel
    - Menggunakan perintah manual : **go test -v -count=1 ./…** (bukan dari tombol Run Test di VSCode)

      Tidak terjadi error seperti ***record not found***

###