package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// you think I'm mad right? No, my first target used to have a lot of Hi-Res original Image. So we definitely need a 20s client timeout. *Note : my first target was... uhm, I don't want to mention her name. I had crush on her...
var clientTwenty *http.Client = &http.Client{Timeout: 20 * time.Second}

func GetPage(url string) ([]byte, string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error on creating http request :\n", err.Error())
		return nil, "", err
	}

	resp, err := clientTwenty.Do(req)
	if err != nil {
		fmt.Println("Error on doing http request :\n", err.Error())
		return nil, "", err
	}
	respByte, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Error on reading response body :\n", err.Error())
		return nil, "", err
	}

	return respByte, string(respByte), nil
}
