package main

import (
	"encoding/base64"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	url = "https://wikin.herokuapp.com/pages.json"
)

type Page struct {
	Id, Url, Title, Body string
}

func authInfo() string {
	user := os.Getenv("WIKIN_USER")
	password := os.Getenv("WIKIN_PASSWORD")
	return base64.StdEncoding.EncodeToString([]byte(user + ":" + password))
}

func get() []byte {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Basic "+authInfo())

	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return byteArray
}

func parse(jsonString []byte) []interface{} {
	js, err := simplejson.NewJson(jsonString)

	if err != nil {
		log.Fatal(err)
	}

	array, err := js.Get("pages").Array()
	if err != nil {
		log.Fatal(err)
	}
	return array
}

func main() {
	jsonString := get()
	pages := parse(jsonString)

	for _, page := range pages {
		fmt.Println(page)
	}
}
