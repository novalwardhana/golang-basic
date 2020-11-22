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
	fmt.Println("-----")
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
	fmt.Println("-----")
}

func prepareQuery() {
	db, err := connect()
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}
	defer db.Close()

	smtp, err := db.Prepare("select * from tb_student where id = ?")
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}

	var student1 = student{}
	smtp.QueryRow("B001").Scan(&student1.name, &student1.id, &student1.grade, &student1.age)
	fmt.Printf("Student B001: %#v\n", student1)

	var student2 = student{}
	smtp.QueryRow("B002").Scan(&student2.age, &student2.grade, &student2.id, &student2.name)
	fmt.Printf("Student B002: %#v\n", student2)

	var student3 = student{}
	smtp.QueryRow("E001").Scan(&student3.age, &student3.grade, &student3.id, &student3.name)
	fmt.Printf("Student E001: %#v\n", student3)

	fmt.Println("-----")
}

func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into tb_student values(?, ?, ?, ?)", "W002", "Ronaldo", 25, 2)
	if err != nil {
		fmt.Printf("Err: %s\n", err.Error())
		return
	}
	fmt.Println("Success added new row")

	_, err = db.Exec("update tb_student set name = ? where id = ?", "Rooney", "W002")
	if err != nil {
		fmt.Printf("Err: %s\n", err.Error())
		return
	}
	fmt.Println("Success update row")

	_, err = db.Exec("delete from tb_student where id= ? ", "W002")
	if err != nil {
		fmt.Printf("Err: %s\n", err.Error())
		return
	}
	fmt.Println("Success delete data")

	fmt.Println("-----")
}

func main() {
	fmt.Println("Part 56 SQL")
	getStudents()
	getStudent()
	prepareQuery()
	sqlExec()
}
