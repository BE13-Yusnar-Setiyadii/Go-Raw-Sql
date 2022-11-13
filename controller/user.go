package controller

import (
	"be13/yusnarsetiyadi/entities"
	"database/sql"
	"fmt"
	"log"
)

func GetAllUser(db *sql.DB) ([]entities.User, error) {
	result, errSelect := db.Query("select id, name, gender, status from users")
	if errSelect != nil {
		return nil, errSelect
	}

	var dataUser []entities.User
	for result.Next() {
		var userrow entities.User
		errScan := result.Scan(&userrow.Id, &userrow.Name, &userrow.Gender, &userrow.Status)
		if errScan != nil {
			return nil, errScan
		}
		dataUser = append(dataUser, userrow)
	}
	return dataUser, nil
}

func InsertDataUser(db *sql.DB, newUser entities.User) {
	var query = "insert into users (id, name, gender, status) values (?,?,?,?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error prepare insert", errPrepare.Error())
	}

	result, errExec := statement.Exec(newUser.Id, newUser.Name, newUser.Gender, newUser.Status)
	if errExec != nil {
		log.Fatal("error exec insert", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("insert berhasil")
		} else {
			fmt.Println("insert gagal")
		}
	}
}

func UpdateDataUser(db *sql.DB, updateUser entities.User) {
	var query = "update users set name = ?, gender = ?, status = ? where id = ?"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error prepare update", errPrepare.Error())
	}

	result, errExec := statement.Exec(updateUser.Name, updateUser.Gender, updateUser.Status, updateUser.Id)
	if errExec != nil {
		log.Fatal("error exec update", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("update berhasil")
		} else {
			fmt.Println("update gagal")
		}
	}
}

func DeleteDataUser(db *sql.DB, deleteUser entities.User) {
	var query = "delete from users where id = ?"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error prepare delete", errPrepare.Error())
	}
	result, errExec := statement.Exec(deleteUser.Id)
	if errExec != nil {
		log.Fatal("error exec delete", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("delete berhasil")
		} else {
			fmt.Println("delete gagal")
		}
	}
}

func ReadByIdDataUser(db *sql.DB, readbyidUser entities.User) {
	result := db.QueryRow(" select id, name, gender, status from users where id = ?", readbyidUser.Id)
	var userrow entities.User
	errScan := result.Scan(&userrow.Id, &userrow.Name, &userrow.Gender, &userrow.Status)
	if errScan != nil {
		fmt.Println("id tidak ditemukan, coba lagi", errScan.Error())
	}
	fmt.Printf("\nid: %d\nnama: %s\ngender: %s\nstatus: %s", userrow.Id, userrow.Name, userrow.Gender, userrow.Status)
}
