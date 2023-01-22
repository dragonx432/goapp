package main

import (
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "your-user"
    password = "your-password"
    dbname   = "your-dbname"
)

func main() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully connected to PostgreSQL database!")

    // Insert data into table
    sqlStatement := `
    INSERT INTO users (name, email)
    VALUES ($1, $2)`
    _, err = db.Exec(sqlStatement, "John Doe", "johndoe@example.com")
    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully inserted data into the users table!")
}
