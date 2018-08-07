package main

import (
	"fmt"
)

// UserID is the unique ID you find on someone's flickr photostream. For example : 6660129401173@N02
const UserID = ""

func main() {
	pageRaw, pageString, err := GetPage(GetMainPageLink(UserID))
	if err != nil {
		fmt.Println("[ERROR] Error on getting the main page raw :\n", err.Error())
	}

	allImageIDs := make([]string, 0)
	OriginalImageURLs := make([]string, 0)
	mainPageLinks := GetAllLinks(pageRaw)
	maxPage := GetMaximumPageIndex(mainPageLinks)

	mainPageImageIDs := FindImageIDs(pageString)
	for _, vmpii := range mainPageImageIDs {
		allImageIDs = append(allImageIDs, vmpii)
	}

	for i := 2; i <= maxPage; i++ {
		currentPageURL := GetPagesLink(UserID, i)
		_, pageString, err := GetPage(currentPageURL)
		if err != nil {
			fmt.Println("[ERROR] Error on getting the next page raw :\n", err.Error())
			continue
		}
		nextPageImageIDs := FindImageIDs(pageString)
		for _, vmpii := range nextPageImageIDs {
			allImageIDs = append(allImageIDs, vmpii)
		}
	}

	for _, imgID := range allImageIDs {
		originalURL := GetOriginalImagePageLink(UserID, imgID)
		originalPage, originalString, err := GetPage(originalURL)
		fmt.Println(originalURL)
		if err != nil {
			fmt.Println("[ERROR] Error on getting the original page raw :\n", err.Error())
			continue
		}
		OriginalImageURL := GetOriginalImageLink(originalPage)
		fmt.Println("-->", OriginalImageURL)
		OriginalImageURLs = append(OriginalImageURLs, OriginalImageURL)
		if 1 > 2 {
			fmt.Println(pageString, originalString)
		}

	}
	for koiu, voiu := range OriginalImageURLs {
		fmt.Println("Downloading Image [", koiu+1, "of", len(OriginalImageURLs), "]")
		err := Download(voiu)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if koiu+1 == len(OriginalImageURLs) {
			fmt.Println("--> Download Finished, enjoy the result :) <--")
		}
	}

}
