package internals

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Id        int64  `json:"id"`
	PCName    string `json:"pc_name"`
	MacAdress string `json:"mac_adress"`
	UserName  string `json:"user_name"`
}

const (
	BaseURL = "http://localhost:8080/users"
)

func Create(args []string) error {
	url := BaseURL + "/create"
	fmt.Println("URL:>", url)

	us := &User{
		PCName:    args[0],
		MacAdress: args[1],
		UserName:  args[2],
	}

	jsonStr, err := json.Marshal(us)
	if err != nil {
		return err
	}

	err = sendRequest("POST", url, jsonStr)
	if err != nil {
		return err
	}

	return nil
}

func Update(args []string) error {
	url := BaseURL + "/update/"
	fmt.Println("URL:>", url)

	if len(args) < 4 {
		fmt.Println("Need to all arguments")
	}

	url = url + args[0]

	us := &User{
		PCName:    args[1],
		MacAdress: args[2],
		UserName:  args[3],
	}

	jsonStr, err := json.Marshal(us)
	if err != nil {
		return err
	}

	err = sendRequest("PUT", url, jsonStr)
	if err != nil {
		return err
	}

	return nil
}

func Delete(args []string) error {
	url := BaseURL + "/delete/"
	fmt.Println("URL:>", url)

	us := &User{}

	url = url + args[0]

	jsonStr, err := json.Marshal(us)
	if err != nil {
		return err
	}

	err = sendRequest("DELETE", url, jsonStr)
	if err != nil {
		return err
	}

	return nil
}

func sendRequest(method, url string, data []byte) error {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
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
