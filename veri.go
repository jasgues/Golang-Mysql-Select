package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {

	/*dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "goo"*/
	
	db, err := sql.Open("mysql", "KULLANICIADI:PAROLA@tcp(SUNUCUADRESİ:3306)/VERİTABANI_ADI")
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Post struct {
	Id    int
	Mail  string
	Pass  string
	Boyut string
}

func getAll() {

	db := dbConn()

	selDB, err := db.Query("SELECT * FROM veri ORDER BY id ASC") //DESC
	if err != nil {
		panic(err.Error())
	}

	post := Post{}
	posts := []Post{}

	for selDB.Next() {

		var id int
		var mail, pass, boyut string

		err = selDB.Scan(&id, &mail, &pass, &boyut)
		if err != nil {
			panic(err.Error())
		}

		post.Id = id
		post.Mail = mail
		post.Pass = pass
		post.Boyut = boyut

		posts = append(posts, post)
	}

	for _, post := range posts {
		fmt.Println(post.Id, post.Mail, post.Pass, post.Boyut)

	}

	defer db.Close()
	fmt.Scanf("h")
} //Getall fonksiyonu
func main() {

	getAll()

}
