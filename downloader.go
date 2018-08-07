package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const filePathPrefix = "result/"

func Download(url string) error {
	fmt.Println("[DOWNLOADER] Downloading", url, ". . . . .")
	fileName := filePathPrefix + getURLFileNameSuffix(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("[DOWNLOADER] Failed to download image : ", err.Error())
		return err
	}
	defer resp.Body.Close()
	imageFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("[DOWNLOADER] Failed to save image : ", err.Error())
		return err
	}
	_, err = io.Copy(imageFile, resp.Body)
	if err != nil {
		fmt.Println("[DOWNLOADER] Failed to copy image : ", err.Error())
		return err
	}
	imageFile.Close()
	return nil
}
func getURLFileNameSuffix(url string) string {
	fileName := ""
	limit := len(url) - 1
	stopper := false
	for !stopper {
		curr := url[limit : limit+1]
		if curr != "/" {
			fileName = curr + fileName
			limit--
		} else {
			stopper = true
		}
	}
	return fileName
}
