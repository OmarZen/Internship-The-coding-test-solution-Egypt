package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Phone  string `json:"phone,omitempty"`
}

type Response struct {
	StatusCode int    `json:"status_code"`
	Data       []User `json:"data"`
}

func main() {
	// Establish a connection to the PostgreSQL database
	db, err := sql.Open("postgres", "host=localhost port=5432 user=your_username password=your_password dbname=your_database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Execute the query to fetch data from the user_table
	rows, err := db.Query("SELECT * FROM public.user_table")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User

	// Iterate over the query result and populate the users slice
	for rows.Next() {
		var user User
		err := rows.Scan(&user.UserID, &user.Name, &user.Age, &user.Phone)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Create a response object with the fetched data
	response := Response{
		StatusCode: 200,
		Data:       users,
	}

	// Convert the response object to JSON
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	// Print the JSON data
	fmt.Println(string(jsonData))
}
