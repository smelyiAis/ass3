package main

import (
    "fmt"
    "sync"
)

// Database represents a database connection, and it's implemented as a Singleton.
type Database struct {
    connection string
}

var instance *Database
var once sync.Once

// GetInstance returns the single instance of the Database.
func GetInstance() *Database {
    once.Do(func() {
        instance = &Database{
            connection: "MyDatabaseConnection",
        }
    })
    return instance
}

// Query performs a database query.
func (db *Database) Query(query string) {
    fmt.Printf("Executing query '%s' on database '%s'\n", query, db.connection)
}

func main() {
    // Get the singleton instance
    dbInstance1 := GetInstance()
    fmt.Println("Database instance 1: ", dbInstance1)

    // Perform a query using the first instance
    dbInstance1.Query("SELECT * FROM table1")

    // Get the singleton instance again
    dbInstance2 := GetInstance()
    fmt.Println("Database instance 2: ", dbInstance2)

    // These two instances are the same
    if dbInstance1 == dbInstance2 {
        fmt.Println("Both instances are the same.")
    }

    // Perform a query using the second instance
    dbInstance2.Query("INSERT INTO table2 (data) VALUES ('some data')")
}
