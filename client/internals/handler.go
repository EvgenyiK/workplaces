package internals

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	BaseURL = "http://localhost:8080/users"
)

func Create() error {

	url := BaseURL + "/create"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{"pc_name":"bpc","mac_adress":"1234g","user_name":"bars"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}
