package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	username = "postgres"
	password = "masbayu28"
	dbName   = "db-go-sql"
)

var (
	db  *sql.DB
	err error
)

type Employee struct {
	ID       int
	FullName string
	Email    string
	Age      int
	Division string
}

func createEmployee() {
	employee := Employee{}

	query := `
		INSERT INTO employees (full_name, email, age, division)
		VALUES ($1, $2, $3, $4)
		RETURNING *
	`

	err = db.QueryRow(query, "Marsha", "marsha@youtube.com", 18, "Idol").
		Scan(&employee.ID, &employee.FullName, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Success to add new employee: %+v\n", employee)
}

func getEmployees() {
	result := []Employee{}

	query := "SELECT * FROM employees"

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		employee := Employee{}

		err := rows.Scan(&employee.ID, &employee.FullName, &employee.Email, &employee.Age, &employee.Division)
		if err != nil {
			panic(err)
		}

		result = append(result, employee)
	}

	fmt.Printf("List of employee: %+v\n", result)
}

func updateEmployee() {
	query := `
		UPDATE employees
		SET full_name=$2, email=$3, age=$4, division=$5
		WHERE id=$1
	`

	res, err := db.Exec(query, 1, "Windah Basudara", "windah@basudara.com", 25, "Finance")
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated data count:", count)
}

func deleteEmployee() {
	query := `
		DELETE from employees
		WHERE id=$1
	`

	res, err := db.Exec(query, 1)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted data count:", count)
}

func main() {
	psqlURL := fmt.Sprintf("host=%s port=%d user=%s password=%s "+
		"dbname=%s sslmode=disable", host, port, username, password, dbName)

	db, err = sql.Open("postgres", psqlURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection to postgresql established")
	getEmployees()
	updateEmployee()
	getEmployees()
	deleteEmployee()
	getEmployees()
	// createEmployee()
}
