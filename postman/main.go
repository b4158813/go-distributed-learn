package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	RegistryRouter = "/services"
	LogRouter      = "/log"
	AppRouter      = "/app"
)

func SendPost(content map[string]string, port, router string) error {
	data, _ := json.Marshal(content)
	resp, err := http.Post(
		"http://127.0.0.1:"+port+router,
		"text/log",
		strings.NewReader(string(data)),
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
	SendPost(map[string]string{
		"servicename": "MyNewService",
		"serviceurl":  "127.0.0.1:4000/log",
	}, "3000", RegistryRouter)
}
