# MINI PROJECT 2 -> Extended Features from Mini Project 1

### Nama aplikasi         : Aplikasi CLI Manajemen Daftar Buku Perpustakaan
### Fitur pada apliaksi	:

1. Semua fitur yang sudah ada pada Mini project 1 dengan update:
    1. Kode buku yang dapat diinputkan oleh users sebelumnya harus bersifat ***unique*** atau, kode buku tidak boleh sama, ketika user menginput kode buku yang sama maka akan tampil informasi bahwa kode buku sudah digunakan.
    2. Buku yang telah diinputkan akan langsung disimpan kedalam file ***JSON***. File json disimpan pada directory ***books*** dengan format nama file **book-{kode_buku}.json**. Buku yang diinputkan bisa lebih dari 1 kali sebelum kembali ke menu utama.
    3. Untuk menampilkan list buku pada perpustakan, semua data buku diload dari file json pada directory books yang telah disimpan tadi.  Ini memungkin kan aplikasi ini dapat penyimpan data buku selamanya.
    4. Dapat menghapus buku yang dipilih serta sekaligus menghapus file json pada directory books.
    5. Dapat mengubah/edit detail buku lalu akan secara otomatis mengupdate file json yang tersimpan.
2. Terdapat fitur baru yaitu **“Print Buku”** dimana pada menu ini users akan diminta untuk memilih buku yang ingin di print, ketika user sudah memilih maka aplikasi dapat melakukan konversi data buku ke dalam sebuah file document ***pdf*** yang disimpan pada directory ***pdf***. Selain itu users juga dapat memilih untuk melakukan print semua data buku daripada memilih satu-satu.

###


