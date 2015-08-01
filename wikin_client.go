package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	url = "https://wikin.herokuapp.com/pages.json"
)

func authInfo() string {
	user := os.Getenv("WIKIN_USER")
	password := os.Getenv("WIKIN_PASSWORD")
	return base64.StdEncoding.EncodeToString([]byte(user + ":" + password))
}

func get() string {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Basic "+authInfo())

	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("connect error\n")
		fmt.Println(err)
		return ""
	}

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error\n")
		fmt.Println(err)
		return ""
	}

	return string(byteArray)
}

func main() {
	get()
}
