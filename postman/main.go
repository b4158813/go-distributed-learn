package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	LogRouter = "/log"
	AppRouter = "/app"
)

func SendPost(content, port, router string) error {
	resp, err := http.Post(
		"http://127.0.0.1:"+port+router,
		"text/log",
		strings.NewReader(content),
	)
	if err != nil {
		fmt.Println("http.Post error: ", err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll error: ", err)
		return err
	}
	fmt.Println(string(body))
	return nil
}

func main() {
	SendPost("12345", "80", LogRouter)
}
