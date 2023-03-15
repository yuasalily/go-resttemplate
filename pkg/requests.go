package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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
	fmt.Println("response: ", response)
	return response

}
