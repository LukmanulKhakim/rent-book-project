package main

import (
	"fmt"
	"project-rent/controller"
	"project-rent/model"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// koneksi ke gorm
func connectGorm() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/latihan_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}

// pembuatan tabel otomatis
func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Book{})

}

func main() {
	var run bool = true
	var menu int
	gconn, err := connectGorm() //panggil
	migrate(gconn)

	//pemanggilan di model dan kontroler
	bukuMDL := model.GormBookModel{gconn}
	bukuCTL := controller.GormBookController{bukuMDL}
	if err != nil {
		fmt.Println("cannot connect to database", err.Error())
	}
	for run {
		fmt.Println("\t--menu--")
		fmt.Println("1. Input produk")
		fmt.Println("2. Update Stok")
		fmt.Println("3. Transaksi")
		fmt.Println("4. Customer")
		fmt.Println("9. Tutup")
		fmt.Println("Masukkan input : ")
		fmt.Scanln(&menu)
		switch menu {
		case 1:
			var tampilanBuku model.Book
			fmt.Print("nama :")
			fmt.Scanln(&tampilanBuku.Nama_buku)
			fmt.Print("harga :")
			fmt.Scanln(&tampilanBuku.Penulis)
			fmt.Print("stok :")
			fmt.Scanln(&tampilanBuku.Rent)

			newProd, err := bukuCTL.Insert(bukuBaru)
			if err != nil {
				fmt.Println("Eror insert", err.Error())
			}
			fmt.Println("Selesai input produk", newProd)
		case 2:
		case 3:
		case 4:
		case 9:
			run = false
			fmt.Println("Tutup")
		}
	}
}
