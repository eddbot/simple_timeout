package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	tt := time.Now()
	simple()
	fmt.Println(time.Since(tt))
}

func simple() {

	timeout := time.After(500 * time.Millisecond)
	ch := make(chan string)
	go func() {
		res, err := slowCall()
		if err != nil {
			panic(err)
		}
		ch <- res
	}()

	select {
	case <-timeout:
		fmt.Println("time out!")
	case hello := <-ch:
		fmt.Println(hello)
	}
}

func slowCall() (string, error) {

	req, err := http.NewRequest(http.MethodGet, "https://poetrydb.org/title/Ozymandias/lines.json", nil)

	if err != nil {
		return "", err

	}
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
