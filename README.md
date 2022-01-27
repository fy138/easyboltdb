# easyboltdb
easy to use boltDB
```go
package main

import (
	"log"

	"github.com/fy138/easyboltdb"
)

func main() {
	//create a db and a table
        // db  file name and  bucket name
	sdb, err := easyboltdb.NewBoltDB("name.db", "user")
	if err != nil {
		log.Fatal(err)
	}
	//save 
	sdb.Update("author", "fy138")
	sdb.Update("today", "2022-01-27")
	//get 
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

```
