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
			fmt.Println("4. Update Profile")
			fmt.Println("5. My Rent")
			fmt.Println("9. Logout")
			fmt.Print("Enter Number : ")
			fmt.Scanln(&menu)

			switch menu {
			case 1:
			case 2:
			case 3:
			case 4:
			case 5:
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
