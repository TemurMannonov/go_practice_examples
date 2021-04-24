package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

// MarshalJSON for NullInt64
func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

// UnmarshalJSON for NullInt64
func (ni *NullInt64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ni.Int64)
	ni.Valid = (err == nil)
	return err
}

// MAIN program starts here
func main() {

	connStr := "user=temur dbname=test password=12345 host=127.0.0.1 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Print(err)
	}

	//sql to insert student information
	statement := "INSERT INTO student(fname, lname) VALUES($1, $2)"
	//prepare statement for sql
	stmt, err := db.Prepare(statement)
	if err != nil {
		fmt.Print(err)
	}
	defer stmt.Close()

	rows, err := db.Query("Select fname, lname from student")
	if err != nil {
		fmt.Print(err)
	}
	defer rows.Close()

	for rows.Next() {

		var fname string
		var lname string
		err := rows.Scan(&fname, &lname)
		if err != nil {
			fmt.Print(err)
		}

		fmt.Printf("%s %s\n", fname, lname)
	}
}
