package main

import (
	"context"
	"crypto/tls"
	"fmt"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

func main() {
	fmt.Println("I'm starting the test")
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
		TLSConfig: &tls.Config{ /*...*/ },
	})
	if err != nil {
		fmt.Println("There was an error creating the client")
	}
	fmt.Println("I'm about to create the client")
	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", ""),
	})
	if err != nil {
		fmt.Println("There was an error creating the connection")
	}
	fmt.Println("I'm about to connect to the db")
	ctx := context.Background()
	db, err1 := c.Database(ctx, "test")
	if err1 != nil {
		fmt.Println("There was an error connecting to the database")
	}
	fmt.Println("I'm about to connect to the collection")
	found, err2 := db.CollectionExists(ctx, "myCollection")
	if found {
		fmt.Println("The collection exists!")
	} else {
		fmt.Println("The collection does NOT exist... you should create it...")
	}
	if err2 != nil {
		fmt.Println("There was an error connecting to the collection")
	}
	fmt.Println("I'm done, going home")
}
