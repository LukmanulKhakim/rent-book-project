package main

import (
	"fmt"
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
func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}

func main() {
	var run bool = true
	var menu int
	gconn, err := connectGorm() //panggil
	migrate(gconn)
	userMDL := model.UserModel{gconn}
	userCTL := controller.UserController{userMDL}
	if err != nil {
		fmt.Println("cannot connect to datavbase", err.Error())
	}
	for run {
		fmt.Println("\t--menu--")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Update Profil")
		fmt.Println("4. Non Aktif User")
		fmt.Println("5. Lihat semua user")
		fmt.Println("9. Tutup")
		fmt.Println("Masukkan input : ")
		fmt.Scanln(&menu)
		switch menu {
		case 1:
			var userBaru model.User
			fmt.Print("nama :")
			fmt.Scanln(&userBaru.Nama)
			fmt.Print("email :")
			fmt.Scanln(&userBaru.Email)
			fmt.Print("pass :")
			fmt.Scanln(&userBaru.Password)
			fmt.Print("addres :")
			fmt.Scanln(&userBaru.Addres)

			newUser, err := userCTL.Register(userBaru)
			if err != nil {
				fmt.Println("Eror insert", err.Error())
			}
			fmt.Println("Selesai input produk", newUser)
		case 2:
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
		case 3:
			var editUser model.User
			var id_user int
			fmt.Print("id kalian :")
			fmt.Scanln(&id_user)
			fmt.Print("edit nama :")
			fmt.Scanln(&editUser.Nama)
			fmt.Print("edit addres :")
			fmt.Scanln(&editUser.Addres)
			fmt.Print("edit email :")
			fmt.Scanln(&editUser.Email)
			fmt.Print("edit password :")
			fmt.Scanln(&editUser.Password)

			UpdateUser, err := userCTL.Update(id_user, editUser)
			if err != nil {
				fmt.Println("Eror update", err.Error())
			}
			fmt.Println("Update profil berhasil", UpdateUser)

		case 4:
		case 5:
			res, err := userCTL.GetAll()
			if err != nil {
				fmt.Println("eror menampilkan user", err.Error())
			}
			fmt.Println(res)
		case 9:
			run = false
			fmt.Println("Tutup")
		}
	}
}
