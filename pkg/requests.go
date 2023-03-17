package pkg

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func Exchange[T any](url string, method string, body any, headers map[string]string) T {
	byteBody, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(byteBody))
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	var response T

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Fatal(err)
	}
	return response
}

func GetForObject[T any](url string) T {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.StatusCode)
	}

	var response T

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Fatal(err)
	}
	return response
}
