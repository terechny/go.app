package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Postm struct {
	Id    int
	Title string
	Text  string
	Image string
	Date  string
}

var database *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:root@/go_test")
	database = db
	// defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func GetPosts() []Postm {

	ConnectDB()

	rows, err := database.Query("select id, title, text, image, date from go_test.post")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	defer rows.Close()
	posts := []Postm{}

	for rows.Next() {
		p := Postm{}
		err := rows.Scan(&p.Id, &p.Title, &p.Text, &p.Image, &p.Date)
		if err != nil {
			fmt.Println(err)
			continue
		}
		posts = append(posts, p)
	}

	return posts
}

func GetPost(id string) Postm {

	ConnectDB()

	sqlStatement := `SELECT * FROM post WHERE id=?;`

	posts := Postm{}

	row := database.QueryRow(sqlStatement, id)

	err := row.Scan(&posts.Id, &posts.Title, &posts.Text, &posts.Image, &posts.Date)

	if err != nil {
		fmt.Println(err)
	}

	return posts
}

func Delete(id string) {

	ConnectDB()

	_, err := database.Exec("DELETE FROM post WHERE `post`.`id` = ?;", id)

	if err != nil {
		fmt.Println(err)
	}
}

func Update(id string, title string, text string) {

	ConnectDB()

	_, errs := database.Exec("UPDATE `post` SET `title` = ?, `text` = ? WHERE `post`.`id` = ?;",
		title, text, id)

	if errs != nil {
		fmt.Println(errs)
	}
}

func Store(title string, text string, filename string) {

	ConnectDB()

	database.Exec("insert into go_test.post (title,text,image) values (?, ?, ?)",
		title, text, filename)
}
