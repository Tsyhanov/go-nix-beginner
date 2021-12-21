package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//task#3 from NIX Education: get posts from Net
func main() {
	fmt.Println("start")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // Log the request body
	bodyString := string(body)
	fmt.Print(bodyString)
}
