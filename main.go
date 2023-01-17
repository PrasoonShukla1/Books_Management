package main

import (
	"Exercise6/books"
	"Exercise6/db"
	"Exercise6/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// ...
func main() {
	db := db.Connecttomysql()
	d := books.Handler{Db: db}
	fmt.Println(d.Create(models.Book{ISBN: 3, Title: "ZopSmart", Author: "Ganesh", Genre: "Fiction", Publication: "Pearson", Year_of_publication: 2019}))
}
