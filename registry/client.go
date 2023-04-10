package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func RegisterService(r Registration) error {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(r)
	if err != nil {
		log.Println("encode error:", err)
		return err
	}
	resp, err := http.Post(ServiceUrl, "application/json", buf)
	if err != nil {
		log.Println("http.post error:", err)
	}
	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("registration failed! response status code: %v", resp.StatusCode)
		return err
	}
	return nil
}
