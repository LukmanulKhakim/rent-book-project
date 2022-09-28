package main

import (
	"fmt"
	"os"
	"os/exec"
	"project-rent/controller"
	"project-rent/model"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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
	db.AutoMigrate(&model.User{})
}

func main() {
	var run bool = true
	var login bool = false
	var next string
	var menu int
	var or string
	gconn, err := connectGorm() //panggil
	migrate(gconn)
	userMDL := model.UserModel{gconn}
	userCTL := controller.UserController{userMDL}
	BookMDL := model.BookModel{gconn}
	BookCTL := controller.BookControl{BookMDL}
	if err != nil {
		fmt.Println("cannot connect to datavbase", err.Error())
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

			case 2:
				Clear()
				fmt.Println("--Login / Regist --")
				fmt.Println("1. Registrasi")
				fmt.Println("2. Login")
				fmt.Println("9. Exit Program")
				fmt.Println("0. Home")
				fmt.Println("Input Number: ")
				fmt.Scanln(&menu)
				switch menu {
				case 1:
					Clear()
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
					Clear()

					var email string
					var password string
					fmt.Print("email :")
					fmt.Scanln(&email)
					fmt.Print("pass :")
					fmt.Scanln(&password)
					_, err := userCTL.Login(email, password)
					if err != nil {
						fmt.Println("Login eror", err.Error())
					}
					fmt.Println("Login berhasil")
					fmt.Println("Enter untuk ke area Member")
					login = true
					fmt.Scanln(&next)
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
			fmt.Println("Selamat Datang")
			fmt.Println("1. Search Book")
			fmt.Println("2. List Book")
			fmt.Println("3. Add Book")
			fmt.Println("4. Edit Book")
			fmt.Println("5. Delete Book")
			fmt.Println("6. Delete Book")
			fmt.Println("9. Logout")
			fmt.Print("Enter Number : ")
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
			case 3:
				var produkBaru model.Book
				fmt.Print("judul :")
				fmt.Scanln(&produkBaru.Judul)
				fmt.Print("deskripsi :")
				fmt.Scanln(&produkBaru.Deskripsi)
				produkBaru.Is_Rent = false
				produkBaru.Is_Deleted = false
				//	produkBaru.Id_User = belum digabung
				BukuBaru, err := BookCTL.Add(produkBaru)
				if err != nil {
					fmt.Println("Eror insert", err.Error())
				}
				fmt.Println("Selesai input produk", BukuBaru)
			case 4:
				var numberbook int
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
				fmt.Scanln(&numberbook)

				var bukuEdit model.Book = res[numberbook-1]
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
				var numberbook int
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
				fmt.Scanln(&numberbook)
				var bukuDel model.Book = res[numberbook-1]
				fmt.Println("--hapus buku--")
				bukuDelres, err := BookCTL.Delete(bukuDel)
				if err != nil {
					fmt.Println("Failed", "Deleting Book Failed", bukuDelres)
					fmt.Println("", err.Error())
				}
				fmt.Println("Success", "Deleting Book Success", bukuDelres)
			case 9:
				var orYes bool = true
				for orYes {
					fmt.Println("Logout ? (y/n)")
					fmt.Scanln(&or)
					if or == "Y" || or == "y" {
						orYes = false
						login = false
						Clear()
					} else if or == "N" || or == "n" {
						orYes = false
						Clear()
					} else {
						orYes = true
					}
				}

			}

		}

	}
}
