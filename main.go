package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HTTPRequest(method, url, body string, headers map[string]string) (respBody string, err error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return respBody, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return respBody, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return respBody, err
	}

	respBody = string(b)
	return respBody, nil
}

func main() {
	fmt.Println("Welcome...")
	mp := make(map[string]string)
	fmt.Println(HTTPRequest("GET", "https://jsonplaceholder.typicode.com/todos/1", "", mp))
}
