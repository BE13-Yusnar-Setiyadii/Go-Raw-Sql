package main

import (
	"be13/yusnarsetiyadi/config"
	"be13/yusnarsetiyadi/controller"
	"be13/yusnarsetiyadi/entities"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConnection := config.ConnectToDB()

	defer dbConnection.Close()

	fmt.Println("MENU:\n1. SELECT\n2. INSERT\n3. UPDATE\n4. DELETE\n5. SELECT BY ID")
	fmt.Println("Masukkan pilihan anda: ")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		{
			dataUser, errGetAll := controller.GetAllUser(dbConnection)
			if errGetAll != nil {
				fmt.Println("error get all user")
			}

			for _, value := range dataUser {
				fmt.Printf("id:%d, name:%s, gender:%s, status:%s\n", value.Id, value.Name, value.Gender, value.Status)
			}
		}
	case 2:
		{
			User := entities.User{}
			newUser := User
			fmt.Println("\nMasukkan id user: ")
			fmt.Scanln(&newUser.Id)
			fmt.Println("Masukkan name user: ")
			fmt.Scanln(&newUser.Name)
			fmt.Println("Masukkan gender user (M/F): ")
			fmt.Scanln(&newUser.Gender)
			fmt.Println("Masukkan status user (active/inactive): ")
			fmt.Scanln(&newUser.Status)

			controller.InsertDataUser(dbConnection, newUser)
		}
	case 3:
		{
			User := entities.User{}
			updateUser := User
			fmt.Println("\nMasukkan id user yg ingin di update: ")
			fmt.Scanln(&updateUser.Id)
			fmt.Println("Masukkan update name: ")
			fmt.Scanln(&updateUser.Name)
			fmt.Println("Masukkan update gender (M/F): ")
			fmt.Scanln(&updateUser.Gender)
			fmt.Println("Masukkan update status (active/inactive): ")
			fmt.Scanln(&updateUser.Status)

			controller.UpdateDataUser(dbConnection, updateUser)
		}
	case 4:
		{
			User := entities.User{}
			deleteUser := User
			fmt.Println("Masukkan id user yg ingin di delete: ")
			fmt.Scanln(&deleteUser.Id)

			controller.DeleteDataUser(dbConnection, deleteUser)
		}
	case 5:
		{
			User := entities.User{}
			readbyidUser := User
			fmt.Println("Masukkan id user: ")
			fmt.Scanln(&readbyidUser.Id)

			controller.ReadByIdDataUser(dbConnection, readbyidUser)
		}
	}
}
