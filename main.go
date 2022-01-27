package main

import (
	"log"
)

func main() {
	sdb, err := NewStoreDB("name.db", "user")
	if err != nil {
		log.Fatal(err)
	}

	sdb.Update("author", "fy138")
	sdb.Update("today", "2022-01-27")

	n, err := sdb.View("author")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("author:", n)

	c, err := sdb.View("today")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("today:", c)

	sdb.Close()
}
