package main

import (
	"fmt"
	"log"

	"github.com/chera-mihiretu/Ethiopian-Market-of-Digital/infrastracture"
	_ "github.com/lib/pq"
)

func main() {
	db, err := infrastracture.NewPostDB()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database is connected Successfully")

	defer db.Close()

}
