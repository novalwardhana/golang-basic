package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/NovalAgungGolang")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getStudents() {
	db, err := connect()
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("Select * from tb_student")
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}
	defer rows.Close()

	var students []student
	for rows.Next() {
		student := student{}
		err := rows.Scan(&student.id, &student.name, &student.grade, &student.age)
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
		}
		students = append(students, student)
	}
	fmt.Println(students)
}

func getStudent() {
	db, err := connect()
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}
	defer db.Close()

	student := student{}
	err = db.QueryRow("select * from tb_student where id = ?", "B001").Scan(&student.id, &student.name, &student.grade, &student.age)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}

	fmt.Printf("Student: %v\n", student)
}

func main() {
	fmt.Println("Part 56 SQL")
	getStudents()
	getStudent()
}
