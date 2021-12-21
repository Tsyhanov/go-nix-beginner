package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func sendrequest(wg *sync.WaitGroup, req string, filename string) {
	defer wg.Done()

	resp, err := http.Get(req)
	if err != nil {
		fmt.Printf("Request Failed: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	err = ioutil.WriteFile(filename, []byte(body), 0644)
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Println(string(body))
}

func sendrequestbuf(wg *sync.WaitGroup, req string, filename string) {
	defer wg.Done()

	resp, err := http.Get(req)
	if err != nil {
		fmt.Printf("Request Failed: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	hfile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Request Failed: %s", err)
	}
	defer hfile.Close()

	// make a write buffer
	writer := bufio.NewWriter(hfile)
	//write bytes to buffer
	bytesWritten, err := writer.Write(body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesWritten)
	// Write memory buffer to disk
	writer.Flush()
}

func main() {
	path := "./storage/posts"
	err := os.MkdirAll(path, 0755)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	params := url.Values{}
	fmt.Println("Start")

	for i := 1; i <= 100; i++ {
		params.Set("Id", strconv.Itoa(i))
		s := "https://jsonplaceholder.typicode.com/posts/" + params.Get("Id")
		wg.Add(1)
		if i < 51 {
			go sendrequest(&wg, s, path+"/"+params.Get("Id"))
		} else {
			go sendrequestbuf(&wg, s, path+"/"+params.Get("Id"))
		}
	}

	wg.Wait()
	fmt.Println("Done")
}
