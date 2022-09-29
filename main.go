package main

import (
	"fmt"
	"os"
	"os/exec"
	"project-rent/controller"
	"project-rent/model"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var UserNow model.User
var BookNow model.Book

func connectGorm() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/project_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil

}

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Rent{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Book{})

}

func main() {
	var run bool = true
	var login bool = false
	var next string
	var menu int
	gconn, err := connectGorm() //panggil
	migrate(gconn)
	userMDL := model.UserModel{gconn}
	userCTL := controller.UserController{userMDL}
	BookMDL := model.BookModel{gconn}
	BookCTL := controller.BookController{BookMDL}
	RentMDL := model.RentModel{gconn}
	RentCTL := controller.RentController{RentMDL}
	if err != nil {
		fmt.Println("cannot connect to database", err.Error())
	}

	for run {
		Clear()

		if !login {
			fmt.Println("--Home--")
			fmt.Println("1. List Book")
			fmt.Println("2. Login/Regist")
			fmt.Println("3. Search Book")
			fmt.Println("9. Exit program")
			fmt.Println("Input Number: ")
			fmt.Scanln(&menu)
			switch menu {
			case 1:
				res, err := BookCTL.GetAll()
				if err != nil {
					fmt.Println("Cant watch list book")
				}
				fmt.Println("List Book")
				fmt.Printf("%4s | %5s| %15s| %15s| %15s|\n", "No", "Id ", "Judul", "Deskripsi", "Status")
				if res != nil {
					i := 1
					var status string
					for _, value := range res {
						if value.Is_Rent {
							status = "Not Available"
						} else {
							status = "Available"
						}
						fmt.Printf("%4d | %5d | %15s | %15s | %15s |\n", i, value.ID, value.Judul, value.Deskripsi, status)
						i++
					}
				} else {
					fmt.Println("\n\t\\tt not found list book")
				}
				fmt.Println("Enter untuk menu lainnya")
				fmt.Scanln(&next)
			case 2:

				fmt.Println("--Login / Regist --")
				fmt.Println("1. Registrasi      ")
				fmt.Println("2. Login           ")
				fmt.Println("9. Exit Program    ")
				fmt.Println("0. Home            ")
				fmt.Println("Input Number:      ")
				fmt.Scanln(&menu)
				switch menu {
				case 1:

					var userBaru model.User
					fmt.Println("--Registrasi--")
					fmt.Print("nama :")
					fmt.Scanln(&userBaru.Nama)
					fmt.Print("email :")
					fmt.Scanln(&userBaru.Email)
					fmt.Print("pass :")
					fmt.Scanln(&userBaru.Password)
					fmt.Print("addres :")
					fmt.Scanln(&userBaru.Addres)

					userBaru.IsDel = 0

					newUser, err := userCTL.Register(userBaru)
					if err != nil {
						fmt.Println("Eror Registrasi", err.Error())
					}
					fmt.Println("Berhasil Registrasi", newUser)
				case 2:
					var email string
					var password string
					fmt.Print("email :")
					fmt.Scanln(&email)
					fmt.Print("pass :")
					fmt.Scanln(&password)
					res, err := userCTL.Login(email, password)
					if err != nil {
						fmt.Println("Login eror", err.Error())

					} else {

						UserNow = res
						login = true
						fmt.Println("Login Berhasil ")
						fmt.Println("Tekan Enter untuk ke Menu Member")
						fmt.Scanln(&next)
					}

				case 9:
					fmt.Println("Terima Kasih")
					run = false
				case 0:
					fmt.Println("Home")
					login = false

				}
			case 3:
			case 9:
				fmt.Println("Terima Kasih")
				run = false
			}
		} else {
			fmt.Println("--Member area--")
			fmt.Println("-Selamat Datang")
			fmt.Println("1. Search Book ")
			fmt.Println("2. List Book   ")
			fmt.Println("3. Add Book    ")
			fmt.Println("4. Edit Book   ")
			fmt.Println("5. Delete Book ")
			fmt.Println("---Rent Area-- ")
			fmt.Println("6. Rent Book   ")
			fmt.Println("7. Return Book ")
			fmt.Println("9. Logout      ")
			fmt.Print("Enter Number :   ")
			fmt.Scanln(&menu)

			switch menu {
			case 1:
			case 2:
				res, err := BookCTL.GetAll()
				if err != nil {
					fmt.Println("Cant watch list book")
				}
				fmt.Println("List Book")
				fmt.Printf("%4s | %5s| %15s| %15s| %15s|\n", "No", "Id ", "Judul", "Deskripsi", "Status")
				if res != nil {
					i := 1
					var status string
					for _, value := range res {
						if value.Is_Rent {
							status = "Not Available"
						} else {
							status = "Available"
						}
						fmt.Printf("%4d | %5d | %15s | %15s | %15s |\n", i, value.ID, value.Judul, value.Deskripsi, status)
						i++
					}
				} else {
					fmt.Println("\n\t\\tt not found list book")
				}
				fmt.Println("Enter untuk menu lainnya")
				fmt.Scanln(&next)
			case 3:
				var produkBaru model.Book

				fmt.Print("judul :")
				fmt.Scanln(&produkBaru.Judul)
				fmt.Print("deskripsi :")
				fmt.Scanln(&produkBaru.Deskripsi)
				produkBaru.Is_Rent = false
				produkBaru.Is_Deleted = false
				produkBaru.ID_User = UserNow.ID
				BukuBaru, err := BookCTL.Add(produkBaru)
				if err != nil {
					fmt.Println("Eror insert", err.Error())
				}
				fmt.Println("Selesai input produk", BukuBaru)
			case 4:
				var number int
				fmt.Println("List my book")

				res, err := BookCTL.GetAll()
				if err != nil {
					fmt.Println("Cant watch my book")
				}
				fmt.Println("List Book")
				fmt.Printf("%4s | %5s| %15s| %15s| %15s|\n", "No", "Id ", "Judul", "Deskripsi", "Status")
				if res != nil {
					i := 1
					var status string
					for _, value := range res {
						if value.Is_Rent {
							status = "Not Available"
						} else {
							status = "Available"
						}
						fmt.Printf("%4d | %5d | %15s | %15s | %15s |\n", i, value.ID, value.Judul, value.Deskripsi, status)
						i++
					}
				} else {
					fmt.Println("\n\t\\tt not found list book")
				}
				fmt.Println("Number book for edit")
				fmt.Scanln(&number)

				var bukuEdit model.Book = res[number-1]
				fmt.Println("Tekan Enter untuk skip")
				fmt.Println("Input Judul baru :")
				fmt.Scanln(&bukuEdit.Judul)
				fmt.Println("Input Deskrispi baru :")
				fmt.Scanln(&bukuEdit.Deskripsi)
				bukuEditres, err := BookCTL.Edit(bukuEdit)
				if err != nil {
					fmt.Println("eror edit")
				}
				fmt.Println("sukses", bukuEditres)
			case 5:
				var number int
				fmt.Println("List my book")

				res, err := BookCTL.GetAll()
				if err != nil {
					fmt.Println("Cant watch my book")
				}
				fmt.Println("List Book")
				fmt.Printf("%4s | %5s| %15s| %15s| %15s|\n", "No", "Id ", "Judul", "Deskripsi", "Status")
				if res != nil {
					i := 1
					var status string
					for _, value := range res {
						if value.Is_Rent {
							status = "Not Available"
						} else {
							status = "Available"
						}
						fmt.Printf("%4d | %5d | %15s | %15s | %15s |\n", i, value.ID, value.Judul, value.Deskripsi, status)
						i++
					}
				} else {
					fmt.Println("\n\t\\tt not found list book")
				}
				fmt.Println("Number book for delete")
				fmt.Scanln(&number)
				var bukuDel model.Book = res[number-1]
				fmt.Println("--hapus buku--")
				bukuDelres, err := BookCTL.Delete(bukuDel)
				if err != nil {
					fmt.Println("Failed", "Deleting Book Failed", bukuDelres)
					fmt.Println("", err.Error())
				}
				fmt.Println("Success", "Deleting Book Success", bukuDelres)
			case 6:

				resNotRent, err := BookCTL.NotRent()
				if err != nil {
					fmt.Println("Cant show list book", err.Error())
				}
				fmt.Println("List Book")
				fmt.Printf("%4s | %5s | %15s | %15s | %15s | %15s |\n", "No", "ID_Buku", "Judul", "Deskrpisi", "Status", "pemilik")

				if resNotRent != nil {
					i := 1
					var status string
					for _, value := range resNotRent {
						if value.Is_Rent {
							status = "Not Available"
						} else {
							status = "Available"
						}
						fmt.Printf("%4d | %5d | %15s | %15s | %15s | %5d| \n", i, value.ID, value.Judul, value.Deskripsi, status, value.ID_User)
						i++
					}
				} else {
					fmt.Println(" Not Found")
				}
				var Number int
				fmt.Println("Tekan Nomor Buku Untuk di Pinjam")
				fmt.Scanln(&Number)
				var BookRent model.Book = resNotRent[Number-1]

				me, err := userCTL.GetIdUser(BookRent.ID_User)
				if err != nil {
					fmt.Println("Failed ")
				}
				var newRent model.Rent
				newRent.Return_book = time.Time{}
				newRent.Books_IsRent = false
				newRent.Books_Nama = me.Nama
				newRent.Books_Email = me.Email
				newRent.Judul_Book = BookRent.Judul
				newRent.Deskripsi_Book = BookRent.Deskripsi
				newRent.ID_Buku = BookRent.ID
				newRent.ID_User = UserNow.ID

				resRentBook, err := RentCTL.AddRent(newRent)

				if err != nil {
					fmt.Println("Eror Rent Book", err.Error())
				} else {

					BookRent.Is_Rent = true
					upRentbook, err := BookCTL.Edit(BookRent)
					if err != nil {
						fmt.Println("Eror Update Rent Book", err.Error())
					} else {
						if upRentbook.ID != 0 {
							fmt.Println("Success rent book : "+resRentBook.Judul_Book, "")
						} else {
							fmt.Println("Error on Update isBorrowedBook, no book updated", err.Error())
						}
					}
				}
				fmt.Println("Enter untuk menu lainnya")
				fmt.Scanln(&next)
			case 7:

				var number int
				res, err := RentCTL.GetUserRent(UserNow.ID)
				if err != nil {
					fmt.Println("Cant List Book Rent", err.Error())
				}

				fmt.Println("List Rent Book")
				fmt.Printf("%4s | %5s | %15s | %15s | %15s |\n", "No", "ID Buku", "Judul", "Deskripsi", "Pemilik")

				if res != nil {
					i := 1
					for _, value := range res {
						fmt.Printf("%5d | %5d | %15s | %15s | %5d |\n", i, value.ID, value.Judul_Book, value.Deskripsi_Book, value.ID_User)
						i++
					}
				} else {
					fmt.Println("Book Not Found")
				}

				fmt.Println("Input Nomor buku untuk dikembalikan ")
				fmt.Scanln(&number)
				var BookReturn model.Rent = res[number-1]
				BookReturn.Books_IsRent = true
				upReturnBook, err := RentCTL.Model.RetRent(BookReturn)
				if err != nil {
					fmt.Println("Cant return book")
				} else {
					BookReturn, err := BookCTL.GetBookId(upReturnBook.ID_Buku)
					if err != nil {
						fmt.Println("Failed")
					} else {
						BookReturn.Is_Rent = false
						upIsRent, err := BookCTL.Edit(BookReturn)
						if err != nil {
							fmt.Println("eror update")
						} else {
							if upIsRent.ID > 0 {
								fmt.Println("Return book"+upReturnBook.Books_Nama, "Sucses")
							} else {
								fmt.Println("eror return book to not rent")
							}
						}
					}

				}

			case 9:
				login = false
				Clear()
			}
		}
	}
}
