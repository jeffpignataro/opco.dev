package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	return sb
}

func post(url string, request string) string {
	method := "POST"

	client := &http.Client{}
	payload := strings.NewReader(request)
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err.Error()
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err.Error()
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}
	return string(body)
}
