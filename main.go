package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	ID       int
	Fullname string
	Age      int
	Batch    string
}

func connect() (*sql.DB, error) {
	// "username-mysql:password-mysql@tcp(localhost)/nama-database"

	db, err := sql.Open("mysql", "root:@tcp(localhost)/goschool")
	// tcp 127.0.0.1:3306

	if err != nil {
		panic(err.Error())
	}

	return db, nil
}

func main() {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// var datas []Student

	// query
	// data, err := db.Query("SELECT * FROM students ORDER BY id DESC LIMIT 4")

	//`
	// SELECT
	// logic
	// WHERE //
	// ORDER BY
	// LIMIT
	//`

	// var id = 1
	// data, err := db.Query("SELECT * FROM students WHERE id = ?", id)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// // for range
	// for data.Next() {
	// 	var dat Student
	// 	if data.Scan(&dat.ID, &dat.Fullname, &dat.Age, &dat.Batch); err != nil {
	// 		fmt.Println(err.Error())
	// 		return
	// 	}
	// 	datas = append(datas, dat)
	// }

	// fmt.Println(datas)

	// method QueryRow() untuk mendapat 1 record saja
	// var data = Student{}

	// var id = 1
	// err = db.QueryRow("SELECT * FROM students WHERE id = ?", id).Scan(&data.ID, &data.Fullname, &data.Age, &data.Batch)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println(data)

	// lebih dari 1

	// INSERT
	// _, err = db.Exec("INSERT INTO students VALUES (?, ?, ?, ?)", 5, "pratama", 22, "excellent")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println("success insert to database")

	// INSERT MORE THAN 1 data
	// _, err = db.Exec("INSERT INTO students VALUES (6,'rendra',23,'excellent'), (7,'rendra',23,'excellent'), (8,'rendra',23,'excellent')")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println("success insert 3 data to database")

	// UPDATE
	// _, err = db.Exec("UPDATE students SET age = ?, fullname = ? WHERE id = ?", 100, "jaya-jaya", 2)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println("success update to database")

	// DELETE
	// _, err = db.Exec("DELETE FROM students WHERE id = 2")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println("success delete to database")
}
