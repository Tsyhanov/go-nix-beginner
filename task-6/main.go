package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

//for a first time. we can create it from management tool. Let's try to do it from here :)
func create_db(name string) {
	//connect to mysql
	db, err := sql.Open("mysql", "root:weak_password@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//create db
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE " + name)
	if err != nil {
		panic(err)
	}
	//create tables
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS posts ( user_id integer, id integer, title varchar(125), body varchar(255) )")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS comments ( post_id integer, id integer, name varchar(125), email varchar(50), body varchar(255) )")
	if err != nil {
		panic(err)
	}
}

type Posts []struct {
	Userid int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Comments []struct {
	PostId int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

//insert single post structure into db and request comments for this post
func AddPostToDb(wgposts *sync.WaitGroup, userid int, id int, title string, body string) {
	defer wgposts.Done()
	//connect to mysql
	db, err := sql.Open("mysql", "root:weak_password@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//insert posts
	result, err := db.Exec("insert into posts (user_id, id, title, body) values (?, ?, ?, ?)",
		userid, id, title, body)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())

	//get comments for postID
	GetComments(strconv.Itoa(id))
}

//insert single comment struct into db
func AddCommentToDb(wgcomments *sync.WaitGroup, postid int, id int, name string, email string, body string) {
	defer wgcomments.Done()
	//connect to mysql
	db, err := sql.Open("mysql", "root:weak_password@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//insert comment to db here
	result, err := db.Exec("insert into comments (post_id, id, name, email, body) values (?, ?, ?, ?, ?)",
		postid, id, name, email, body)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())
	fmt.Println("AddCommentToDb for postid" + strconv.Itoa(postid) + ": " + strconv.Itoa(id))
}

//get comments for PostId and start routines to insert it into db
func GetComments(postid string) {
	req := "https://jsonplaceholder.typicode.com/comments?postId=" + postid

	resp, err := http.Get(req)
	if err != nil {
		fmt.Printf("Request Failed: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ReadAll Failed: %s", err)
	}

	comments := Comments{}
	err = json.Unmarshal(body, &comments)
	if err != nil {
		log.Printf("Comments unmarshaling failed: %s", err)
		return
	}

	//create subroutines to insert comments into db
	var wgcomments sync.WaitGroup
	for _, value := range comments {
		wgcomments.Add(1)
		go AddCommentToDb(&wgcomments, value.PostId, value.ID, value.Name, value.Email, value.Body)
	}
	wgcomments.Wait()
}

//Main
func main() {
	fmt.Println("Start")
	create_db("testdb")

	//get posts for userId=7
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts?userId=7")
	if err != nil {
		log.Printf("Request Failed: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ReadAll Failed: %s", err)
	}
	posts := Posts{}
	err = json.Unmarshal(body, &posts)
	if err != nil {
		log.Printf("Posts unmarshaling failed: %s", err)
		return
	}

	//create routines to insert posts into db
	var wgposts sync.WaitGroup
	for _, value := range posts {
		wgposts.Add(1)
		go AddPostToDb(&wgposts, value.Userid, value.ID, value.Title, value.Body)
	}

	wgposts.Wait()
	fmt.Println("Done")
}
