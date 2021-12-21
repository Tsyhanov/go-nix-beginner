package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

func sendrequest(wg *sync.WaitGroup, userid int) {
	defer wg.Done()

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(userid))
	if err != nil {
		fmt.Printf("Request Failed: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if userid < 6 {
		fmt.Println(string(body))
	}
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("Start")

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go sendrequest(&wg, i)
	}
	wg.Wait()
	fmt.Println("Done")
}
